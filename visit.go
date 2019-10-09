package main

import "github.com/antlr/antlr4/runtime/Go/antlr"

import "./pparser"
import "./types"

import "fmt"
import "strconv"
import "math/big"

func visitExprListBool(actx antlr.ParserRuleContext, gv *types.GVal, v *types.PVal) bool {
	ctx := actx.(*parser.ExprListContext)
	gv.Prc = ctx
	for _, child := range ctx.AllExpr() {
		v2 := visitExpr(child, gv)
		if v.Equal(v2) {
			return true
		}
	}
	return false
}

func visitGeneric(actx antlr.ParserRuleContext, gv *types.GVal, args ...interface{}) interface{} {
	var rt interface{}
	gv.Prc = actx
	switch ctx := actx.(type) {

	case *parser.SimpleStatementContext:

		switch v := ctx.GetChild(0).(type) {
		case *parser.CompoundSetContext:
			visitGeneric(v, gv)
		case *parser.SimpleSetContext:
			visitGeneric(v, gv)
		case antlr.TerminalNode:
			if v.GetSymbol().GetTokenType() == parser.PinnParserID {
				opString := ctx.GetChildren()[len(ctx.GetChildren())-1].(antlr.TerminalNode).GetText()
				switch opString {
				case "++", "--":
					var result *types.PVal
					v2 := gv.GetPv(ctx.ID().GetText())
					if ctx.Expr() != nil {
						indexE := visitExpr(ctx.Expr(), gv)
						index := types.FixIndex(indexE.GetInterface())
						lhs := types.NewScalarPval(v2.Get(index))
						result = doOp(ctx, lhs, nil, opString)
						v2.Set(index, result.GetInterface())
						return nil

					}
					result = doOp(ctx, v2, nil, opString)
					gv.PutPv(ctx.ID().GetText(), result)
					return nil
				default:
					panic(types.ErrCase)
				}

			}
		}

	case *parser.ExprListContext:
		rt := make([]*types.PVal, 0, 10)
		for _, child := range ctx.AllExpr() {
			v := visitExpr(child, gv)
			rt = append(rt, v)
		}
		return rt

	case *parser.CompoundSetContext:
		str := ctx.ID().GetText()
		v := gv.GetPv(str)
		children := ctx.GetChildren()
		op := children[len(children)-3].(antlr.TerminalNode).GetText()
		exprs := ctx.AllExpr()
		if len(exprs) == 2 {
			v2 := visitExpr(exprs[0], gv)
			rhs := visitExpr(exprs[1], gv)
			lhs := types.NewScalarPval(v.Get(types.FixIndex(v2.GetInterface())))
			result := doOp(ctx, lhs, rhs, op)
			v.Set(types.FixIndex(v2.GetInterface()), result.GetInterface())
		} else {
			e := visitExpr(exprs[0], gv)
			result := doOp(ctx, v, e, op)
			v.Set(nil, result.GetInterface())
		}

	case *parser.BlockContext:
		for _, child := range ctx.AllStatement() {
			/*
				valAndType(child)
				valAndType(child.GetPayload())
			*/
			visitStatement(child, gv)
			if !gv.IsPathNormal() {
				break
			}
		}

	case *parser.VarDeclContext:

		var f *types.Fc
		if gv.Fc == nil {
			f = gv.Gfc
		} else {
			f = gv.Fc
		}
		strID := ctx.ID().GetText()
		if _, ok := f.M[strID]; ok {
			panic(types.ErrReDeclare)
		}
		if ctx.CE() != nil {
			e := visitExpr(ctx.Expr(), gv)
			f.M[strID] = e
			return nil
		}
		k := visitGeneric(ctx.Kind(), gv).(types.Kind)
		if ctx.ExprList() == nil {
			f.M[strID] = k.Iv()
		} else {
			ai := visitGeneric(ctx.ExprList(), gv).([]*types.PVal)
			if k.Int == -1 {
				k.Int = len(ai)
			}
			v := k.Iv()
			switch k.GType {
			case types.GScalar:
				if len(ai) != 1 {
					panic(types.ErrParamLength)
				}
				v.Set(nil, ai[0].GetInterface())
			case types.GArray:
				if k.Int != len(ai) {
					panic(types.ErrParamLength)
				}
				for key, iv := range ai {
					v.Set(key, iv.GetInterface())
				}
			case types.GSlice, types.GMap:
				if len(ai) != 1 {
					panic(types.ErrParamLength)
				}
				v = ai[0]
				k.EqualsError(v.GetKind())
			default:
				panic(types.ErrCase)
			}
			f.M[strID] = v
		}

	case *parser.FvarDeclContext:
		strID := ctx.ID().GetText()
/*
		if _, ok := gv.Fc.M[strID]; ok {
			panic(types.ErrReDeclare)
		}

		gv.Fc.ParamList = append(gv.Fc.ParamList, strID)

*/
		k := visitGeneric(ctx.Kind(), gv).(types.Kind)
		return types.FKind{k, strID}

		
	case *parser.KindContext:

		strTYPE := ctx.TYPES().GetText()
		t := types.MLitToType[strTYPE]
		if ctx.LSQUARE() == nil {
			return types.NewKind(types.GScalar, t, 1)
		} else {
			if ctx.Expr() != nil {
				v := visitExpr(ctx.Expr(), gv)
				x := int(v.GetInt())
				if x == -1 {
					panic(types.ErrRange)
				}
				return types.NewKind(types.GArray, t, x)
			} else if ctx.MAP() != nil {
				return types.NewKind(types.GMap, t, 0)
			} else if ctx.SLICE() != nil {
				return types.NewKind(types.GSlice, t, 0)
			} else if ctx.FILL() != nil {
				return types.NewKind(types.GArray, t, -1)
			} else {
				panic(types.ErrCase)
			}
		}

	case *parser.ReturnStatementContext:
		gv.Fc.Path = types.PExiting
		i := ctx.Expr()
		if i != nil {
			gv.Fc.Rt = visitExpr(i, gv)
		}
	case *parser.SimpleSetContext:
		str := ctx.ID().GetText()
		v := gv.GetPv(str)
		exprs := ctx.AllExpr()
		if len(exprs) == 2 {
			v2 := visitExpr(exprs[0], gv)
			e := visitExpr(exprs[1], gv)
			v.Set(types.FixIndex(v2.GetInterface()), e.GetInterface())
		} else {
			e := visitExpr(exprs[0], gv)
			if e.GetKind().GType == types.GScalar {
				v.Set(nil, e.GetInterface())
			} else {
				gv.PutPv(str, e)
			}
		}

	case *parser.IndexExprContext:
		var rt *types.PVal
		v := gv.GetPv(ctx.ID().GetText())
		if ctx.COLON() != nil {
			lhsv := 0
			if e := ctx.GetFirst(); e != nil {
				lhsv = int(visitExpr(e, gv).GetInt())
			}
			rhsv := v.GetKind().Int
			if e := ctx.GetSecond(); e != nil {
				rhsv = int(visitExpr(e, gv).GetInt())
			}
			rt = v.Slice(lhsv, rhsv)
		} else {
			e := visitExpr(ctx.Expr(0), gv)
			rt = types.NewScalarPval(v.Get(types.FixIndex(e.GetInterface())))
		}
		return rt

	default:
		panic(types.ErrCase)
	}
	return rt

}

func visitFuncExpr(actx antlr.ParserRuleContext, gv *types.GVal) *types.PVal {
	ctx := actx.(*parser.FuncExprContext)
	gv.Prc = ctx
	var rt *types.PVal
	switch ctx.GetStart().GetText() {
	case "len":
		str := ctx.ID().GetText()
		v := gv.GetPv(str)
		if v.GetKind().GType == types.GScalar {
			panic(types.ErrWrongType)
		}
		rt = types.NewScalarPval(int64(v.GetKind().Int))
	case "strLen":
		e := visitExpr(ctx.Expr(), gv)
		rt = types.NewScalarPval(int64(len(e.GetString())))
	case "stringValue":
		e := visitExpr(ctx.Expr(), gv)
		rt = types.NewScalarPval(e.String())
	case "print":
		s := visitGeneric(ctx.ExprList(), gv).([]*types.PVal)
		if len(s) == 1 {

			globalPrinter.print(fmt.Sprint(s[0]))
		} else {
			globalPrinter.print(fmt.Sprint(s))
		}
		globalPrinter.print("\n")
	case "printB":
		e := visitExpr(ctx.Expr(), gv)
		switch e.GetKind().VType {
		case types.TInt:
			globalPrinter.print(fmt.Sprintf("%b", e.GetInt()))
		case types.TBig:
			globalPrinter.print(e.GetBig().Text(2))
		default:
			panic(types.ErrCase)
		}
		globalPrinter.print("\n")
	case "printH":
		e := visitExpr(ctx.Expr(), gv)
		switch e.GetKind().VType {
		case types.TInt:
			globalPrinter.print(fmt.Sprintf("%x", e.GetInt()))
		case types.TBig:
			globalPrinter.print(e.GetBig().Text(16))
		case types.TBigf:
			str := floatToHex(e.GetBigf().Text('p', 0))
			globalPrinter.print(str)
		default:
			panic(types.ErrCase)
		}
		globalPrinter.print("\n")
	case "delete":
		e := visitExpr(ctx.Expr(), gv)
		v := gv.GetPv(ctx.ID().GetText())
		v.DeleteKey(e.GetString())
	default:
		panic(types.ErrCase)
	}
	return rt
}

func visitExpr(actx antlr.ParserRuleContext, gv *types.GVal) *types.PVal {
	ctx := actx.(*parser.ExprContext)
	gv.Prc = ctx
	var rt *types.PVal
	if ctx.IndexExpr() != nil {
		return visitGeneric(ctx.IndexExpr(), gv).(*types.PVal)
	}
	if ctx.FuncExpr() != nil {
		return visitFuncExpr(ctx.FuncExpr(), gv)
	}
	schildren, _ := symTokens(ctx)
	if ctx.GetChildCount() == 1 {
		tok := schildren[0]
		switch tok.GetTokenType() {
		case parser.PinnParserID:
			v := gv.GetPv(tok.GetText())
			if v.GetKind().GType == types.GMap || v.GetKind().GType == types.GSlice || v.GetKind().GType == types.GScalar {
				rt = v
			} else {
				rt = v.Clone()
			}
		case parser.PinnParserINT:
			x, _ := strconv.ParseInt(tok.GetText(), 0, 64)
			rt = types.NewScalarPval(x)
		case parser.PinnParserFLOAT:
			str := floatToDec(tok.GetText())
			bf, _, _ := new(big.Float).Parse(str, 0)
			rt = types.NewScalarPval(bf)
		case parser.PinnParserBOOL:
			switch tok.GetText() {
			case "true":
				rt = types.NewScalarPval(true)
			case "false":
				rt = types.NewScalarPval(false)
			default:
				panic(types.ErrCase)
			}
		case parser.PinnParserSTRING:
			str := tok.GetText()
			rt = types.NewScalarPval(str[1 : len(str)-1])
		case parser.PinnParserCHAR:
			str := tok.GetText()
			str = str[1 : len(str)-1]

			for _, rune := range str {
				rt = types.NewScalarPval(rune)
			}
		case parser.PinnParserIOTA:
			rt = types.NewScalarPval(nil)
		default:
			panic(types.ErrCase)
		}
		return rt
	}
	first := ctx.GetChild(0)
	switch v := first.(type) {
	case antlr.TerminalNode:
		id := v.GetText()
		switch id {
		case "(":
			if ctx.COLON() == nil {
				rt = visitExpr(ctx.Expr(0), gv)
			} else {
				var lhsv int64 = 0
				if e := ctx.GetFirstExpr(); e != nil {
					lhsv = visitExpr(e, gv).GetInt()
				}
				rhsv := visitExpr(ctx.GetSecondExpr(), gv).GetInt()
				if rhsv-lhsv < 0 {
					panic(types.ErrRange)
				}
				rt := types.NewPval(types.GArray, types.TInt, int(rhsv-lhsv))
				for x := lhsv; x < rhsv; x += 1 {
					rt.Set(int(x-lhsv), x)
				}
				return rt
			}
		case "-":
			e := visitExpr(ctx.Expr(0), gv)
			switch e.GetKind().VType {
			case types.TInt:
				rt = types.NewScalarPval(-e.GetInt())
			case types.TBig:
				rt = types.NewScalarPval(new(big.Int).Neg(e.GetBig()))
			case types.TBigf:
				rt = types.NewScalarPval(new(big.Float).Neg(e.GetBigf()))
			default:
				panic(types.ErrCase)
			}
		case "!":
			e := visitExpr(ctx.Expr(0), gv)
			rt = types.NewScalarPval(!e.GetBool())
		case "^":
			e := visitExpr(ctx.Expr(0), gv)
			switch e.GetKind().VType {
			case types.TInt:
				rt = types.NewScalarPval(^e.GetInt())
			case types.TBig:
				rt = types.NewScalarPval(new(big.Int).Not(e.GetBig()))
			default:
				panic(types.ErrCase)
			}

		default:
			id = ctx.ID().GetText()
			brace := ctx.GetChild(1).(antlr.TerminalNode).GetSymbol().GetTokenType()
			switch brace {

			case parser.PinnParserLPAREN:
				s := make([]*types.PVal, 0)
				if eL := ctx.ExprList(); eL != nil {
					s = visitGeneric(eL, gv).([]*types.PVal)
				}
				rt = visitFunction(gv.Fkmap[id].FunctionContext, s, id, gv)
			case parser.PinnParserLSQUARE:
				/*
				   valAndType(ctx.Expr(0).GetStart().GetTokenIndex())
				   valAndType(ctx.COLON().GetSymbol().GetTokenIndex())
				   valAndType(ctx.Expr(1).GetStart().GetTokenIndex())
				*/

			default:
				panic(types.ErrCase)
			}
		}
	case *parser.ExprContext:
		op := ctx.GetChild(1).(antlr.TerminalNode).GetText()
		lhs := visitExpr(v, gv)
		if op == "?" {
			b := lhs.GetBool()
			if b {
				return visitExpr(ctx.Expr(1), gv)
			}
			return visitExpr(ctx.Expr(2), gv)
		}
		if op == "&&" {
			lhsv := lhs.GetBool()
			if !lhsv {
				rt = types.NewScalarPval(false)
				break
			}
			rhs := visitExpr(ctx.Expr(1), gv)
			return types.NewScalarPval(rhs.GetBool())
		}
		if op == "||" {
			lhsv := lhs.GetBool()
			if lhsv {
				rt = types.NewScalarPval(true)
				break
			}
			rhs := visitExpr(ctx.Expr(1), gv)
			return types.NewScalarPval(rhs.GetBool())
		}

		rhs := visitExpr(ctx.Expr(1), gv)
		rt = doOp(ctx, lhs, rhs, op)
	}

	return rt
}

func visitFoStatement(actx antlr.ParserRuleContext, gv *types.GVal) {
	ctx := actx.(*parser.FoStatementContext)
	gv.Prc = ctx
	if ctx.RANGE() != nil {
		ids := ctx.AllID()
		valueOnly := len(ids) == 1
		var value *types.PVal
		var key *types.PVal
		if valueOnly {
			value = gv.GetPv(ids[0].GetText())
		} else {
			value = gv.GetPv(ids[1].GetText())
		}
		if !valueOnly {
			key = gv.GetPv(ids[0].GetText())
		}
		ranger := visitExpr(ctx.Expr(), gv)
		switch ranger.GetKind().GType {
		case types.GArray, types.GSlice:
			for x := 0; x < ranger.GetKind().Int; x++ {

				if !valueOnly {
					key.Set(nil, int64(x))
				}
				value.Set(nil, ranger.Get(x))
				visitGeneric(ctx.Block(), gv)
				if gv.PathEndBlock() {
					return
				}
			}
		case types.GMap:
			keys := ranger.GetKeys()
			for _, str := range keys {
				if !valueOnly {
					key.Set(nil, str)
				}
				value.Set(nil, ranger.Get(str))
				visitGeneric(ctx.Block(), gv)
				if gv.PathEndBlock() {
					return
				}
			}
		default:
			panic(types.ErrCase)
		}
		return
	}
	if n := ctx.VarDecl(); n != nil {
		visitGeneric(n, gv)
	} else if n := ctx.GetFss(); n != nil {
		visitGeneric(n, gv)
	}
	v := visitExpr(ctx.Expr(), gv)
	for v.GetBool() {
		visitGeneric(ctx.Block(), gv)
		if gv.PathEndBlock() {
			return
		}
		visitGeneric(ctx.GetSss(), gv)
		v = visitExpr(ctx.Expr(), gv)
	}
}

func visitWhStatement(actx antlr.ParserRuleContext, gv *types.GVal) {
	ctx := actx.(*parser.WhStatementContext)
	gv.Prc = ctx
	v := visitExpr(ctx.Expr(), gv)
	for v.GetBool() {
		visitGeneric(ctx.Block(), gv)
		if gv.PathEndBlock() {
			return
		}
		v = visitExpr(ctx.Expr(), gv)
	}
}

func visitCaseStatement(actx antlr.ParserRuleContext, gv *types.GVal, v *types.PVal) bool {
	ctx := actx.(*parser.CaseStatementContext)
	gv.Prc = ctx
	if gv.Fc.Path == types.PFallthrough || visitExprListBool(ctx.ExprList(), gv, v) {
		if gv.Fc.Path == types.PFallthrough {
			gv.Fc.Path = types.PNormal
		}
	loop:
		for _, child := range ctx.AllStatement() {
			visitStatement(child, gv)
			switch gv.Fc.Path {
			case types.PContinue, types.PExiting, types.PFallthrough:
				break loop
			case types.PBreak:
				gv.Fc.Path = types.PNormal
				break loop
			}
		}
		return true
	}
	return false
}

func visitSwitchStatement(actx antlr.ParserRuleContext, gv *types.GVal) {
	ctx := actx.(*parser.SwitchStatementContext)
	gv.Prc = ctx
	v := visitExpr(ctx.Expr(), gv)
	broken := false
	for _, child := range ctx.AllCaseStatement() {
		if visitCaseStatement(child, gv, v) {
			broken = true
			if gv.Fc.Path == types.PFallthrough {
				broken = false
				continue
			}
			break
		}
	}
	if !broken {
		if gv.Fc.Path == types.PFallthrough {
			gv.Fc.Path = types.PNormal
		}
	loop:
		for _, child := range ctx.AllStatement() {
			visitStatement(child, gv)
			switch gv.Fc.Path {
			case types.PContinue, types.PExiting:
				break loop
			case types.PBreak:
				gv.Fc.Path = types.PNormal
				break loop
			}
		}
	}
}

func visitIfStatement(actx antlr.ParserRuleContext, gv *types.GVal) {
	ctx := actx.(*parser.IfStatementContext)
	gv.Prc = ctx
	v := visitExpr(ctx.Expr(), gv)
	if v.GetBool() {
		visitGeneric(ctx.AllBlock()[0], gv)
	} else {
		if ctx.Block(1) != nil {
			visitGeneric(ctx.Block(1), gv)
		}
	}
}

func visitStatement(actx antlr.ParserRuleContext, gv *types.GVal) {
	ctx := actx.(*parser.StatementContext)
	gv.Prc = ctx
	switch v := ctx.GetChild(0).(type) {
	case *parser.ExprContext:
		visitExpr(v, gv)
	case *parser.IfStatementContext:
		visitIfStatement(v, gv)
	case *parser.WhStatementContext:
		visitWhStatement(v, gv)
	case *parser.FoStatementContext:
		visitFoStatement(v, gv)
	case *parser.SwitchStatementContext:
		visitSwitchStatement(v, gv)
	case antlr.TerminalNode:
		switch v.GetText() {
		case "break":
			gv.Fc.Path = types.PBreak
		case "continue":
			gv.Fc.Path = types.PContinue
		case "fallthrough":
			gv.Fc.Path = types.PFallthrough
		case "debug":
			dbg(gv)
		case ";":
		default:
			panic(types.ErrCase)
		}

	default:
		visitGeneric(v.(antlr.ParserRuleContext), gv)
	}
}

func visitFunction(actx antlr.ParserRuleContext, s []*types.PVal, str string, gv *types.GVal) *types.PVal {
	ctx := actx.(*parser.FunctionContext)
	gv.Prc = ctx
	gv.PushFrame()
	oldfc := gv.Fc
	gv.Fc = types.NewFc()
	fh := gv.Fkmap[str]
	if len(s) != len(fh.S) {
		panic(types.ErrParamLength)
	}
	for k, v := range s {
		v.GetKind().EqualsError(fh.S[k].Kind)
		gv.Fc.M[fh.S[k].S] = v
	}
	visitGeneric(ctx.Block(), gv)
	if gv.Fc.Rt.Invalid() {
		if fh.Kind != (types.Kind{}) {
			panic(types.ErrWrongType)
		}
	} else {
		gv.Fc.Rt.GetKind().EqualsError(fh.Kind)
	}
	rt := gv.Fc.Rt
	gv.Fc = oldfc
	return rt
}

func visitFile(actx antlr.ParserRuleContext, b bool) {
	ctx := actx.(*parser.FileContext)
	reset()
	gv := &types.GVal{}
	defer func() {
		if recover() != nil {
			if b {
				dbg(gv)
			}
			panic("pVisitFile")
		}
	}()

	gv.Fkmap = make(map[string]types.Fheader, 20)
	gv.Gfc = types.NewFc()
	gv.Prc = ctx

	for _, child := range ctx.AllFunction() {
		v := child
		visitFunctionHeader(v, gv)
	}
	for _, child := range ctx.AllVarDecl() {
		v := child
		visitGeneric(v, gv)
	}
	visitFunction(gv.Fkmap["main"].FunctionContext, make([]*types.PVal, 0), "main", gv)
	if *bDebug {
		dbg(gv)
	}
}

func visitFunctionHeader(actx antlr.ParserRuleContext, gv *types.GVal) {
	ctx := actx.(*parser.FunctionContext)
	gv.Prc = ctx
	if _, ok := gv.Fkmap[ctx.ID().GetText()]; ok {
		panic(types.ErrReDeclare)
	}
	fh := types.Fheader{FunctionContext: ctx}

	child := ctx.Kind()
	if child != nil {
		fh.Kind = visitGeneric(child, gv).(types.Kind)
	}

	for _, child := range ctx.AllFvarDecl() {
		fh.S = append(fh.S, visitGeneric(child, gv).(types.FKind))
	}
	gv.Fkmap[ctx.ID().GetText()] = fh
}
