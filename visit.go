package main

import "github.com/antlr/antlr4/runtime/Go/antlr"

import "./pparser"
import "./types"

import "fmt"
import "strconv"
import "math/big"

func visitGeneric(actx antlr.ParserRuleContext, gv *types.GVal, args ...interface{}) interface{} {
	var rt interface{}
	oldPrc := gv.Prc
	gv.Prc = actx
	switch ctx := actx.(type) {
	case *parser.CaseStatementContext:
		v := args[0].(*types.PVal)

		if gv.Fc.Path == types.PFallthrough || visitGeneric(ctx.ExprList(), gv, v).(bool) {
			rt = true
			if gv.Fc.Path == types.PFallthrough {
				gv.Fc.Path = types.PNormal
			}
		cloop:
			for _, child := range ctx.AllStatement() {
				visitGeneric(child, gv)
				switch gv.Fc.Path {
				case types.PContinue, types.PExiting, types.PFallthrough:
					break cloop
				case types.PBreak:
					gv.Fc.Path = types.PNormal
					break cloop
				}
			}
		} else {
			rt = false
		}

	case *parser.SwitchStatementContext:
		v := visitGeneric(ctx.Expr(), gv).(*types.PVal)
		broken := false
		for _, child := range ctx.AllCaseStatement() {
			if visitGeneric(child, gv, v).(bool) {
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
				visitGeneric(child, gv)
				switch gv.Fc.Path {
				case types.PContinue, types.PExiting:
					break loop
				case types.PBreak:
					gv.Fc.Path = types.PNormal
					break loop
				}
			}
		}

	case *parser.StatementContext:
		switch v := ctx.GetChild(0).(type) {
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
	case *parser.RepeatStatementContext:
		b := true
		for b {
			visitGeneric(ctx.Block(), gv)
			if gv.PathEndBlock() {
				break
			}
			v := visitGeneric(ctx.Expr(), gv).(*types.PVal)
			b = v.GetBool()
		}

	case *parser.WhStatementContext:
		v := visitGeneric(ctx.Expr(), gv).(*types.PVal)
		for v.GetBool() {
			visitGeneric(ctx.Block(), gv)
			if gv.PathEndBlock() {
				break
			}
			v = visitGeneric(ctx.Expr(), gv).(*types.PVal)
		}
	case *parser.FoStatementContext:
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
			ranger := visitGeneric(ctx.Expr(), gv).(*types.PVal)
			switch ranger.GetKind().GType {
			case types.GArray, types.GSlice:
				for x := 0; x < ranger.GetKind().Int; x++ {

					if !valueOnly {
						key.Set(nil, int64(x))
					}
					value.Set(nil, ranger.Get(x))
					visitGeneric(ctx.Block(), gv)
					if gv.PathEndBlock() {
						break
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
						break
					}
				}
			default:
				panic(types.ErrCase)
			}
		} else {
			if n := ctx.VarDecl(); n != nil {
				visitGeneric(n, gv)
			} else if n := ctx.GetFss(); n != nil {
				visitGeneric(n, gv)
			}
			v := visitGeneric(ctx.Expr(), gv).(*types.PVal)
			for v.GetBool() {
				visitGeneric(ctx.Block(), gv)
				if gv.PathEndBlock() {
					break
				}
				visitGeneric(ctx.GetSss(), gv)
				v = visitGeneric(ctx.Expr(), gv).(*types.PVal)
			}
		}

	case *parser.IfStatementContext:
		v := visitGeneric(ctx.Expr(), gv).(*types.PVal)
		if v.GetBool() {
			visitGeneric(ctx.AllBlock()[0], gv)
		} else {
			if ctx.Block(1) != nil {
				visitGeneric(ctx.Block(1), gv)
			} else if ctx.IfStatement() != nil {
				visitGeneric(ctx.IfStatement(), gv)
			}
		}

	case *parser.FuncExprContext:

		switch ctx.GetStart().GetText() {
		case "len":
			str := ctx.ID().GetText()
			v := gv.GetPv(str)
			if v.GetKind().GType == types.GScalar {
				panic(types.ErrWrongType)
			}
			rt = types.NewScalarPval(int64(v.GetKind().Int))
		case "strLen":
			e := visitGeneric(ctx.Expr(), gv).(*types.PVal)
			rt = types.NewScalarPval(int64(len(e.GetString())))
		case "stringValue":
			e := visitGeneric(ctx.Expr(), gv).(*types.PVal)
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
			e := visitGeneric(ctx.Expr(), gv).(*types.PVal)
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
			e := visitGeneric(ctx.Expr(), gv).(*types.PVal)
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
			e := visitGeneric(ctx.Expr(), gv).(*types.PVal)
			v := gv.GetPv(ctx.ID().GetText())
			v.DeleteKey(e.GetString())
		default:
			panic(types.ErrCase)
		}

	case *parser.ExprContext:
		if ctx.IndexExpr() != nil {
			rt = visitGeneric(ctx.IndexExpr(), gv)
			break
		}
		if ctx.FuncExpr() != nil {
			rt = visitGeneric(ctx.FuncExpr(), gv)
			break
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
			break
		}
		first := ctx.GetChild(0)
		switch v := first.(type) {
		case antlr.TerminalNode:
			id := v.GetText()
			switch id {
			case "(":
				if ctx.COLON() == nil {
					rt = visitGeneric(ctx.Expr(0), gv).(*types.PVal)
				} else {
					var lhsv int64 = 0
					if e := ctx.GetFirstExpr(); e != nil {
						lhsv = visitGeneric(e, gv).(*types.PVal).GetInt()
					}
					rhsv := visitGeneric(ctx.GetSecondExpr(), gv).(*types.PVal).GetInt()
					if rhsv-lhsv < 0 {
						panic(types.ErrRange)
					}
					rt = types.NewPval(types.GArray, types.TInt, int(rhsv-lhsv))
					for x := lhsv; x < rhsv; x += 1 {
						rt.(*types.PVal).Set(int(x-lhsv), x)
					}
				}
			case "+":
				rt = visitGeneric(ctx.Expr(0), gv)
			case "-":
				e := visitGeneric(ctx.Expr(0), gv).(*types.PVal)
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
				e := visitGeneric(ctx.Expr(0), gv).(*types.PVal)
				rt = types.NewScalarPval(!e.GetBool())
			case "^":
				e := visitGeneric(ctx.Expr(0), gv).(*types.PVal)
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
				default:
					panic(types.ErrCase)
				}
			}
		case *parser.ExprContext:
			op := ctx.GetChild(1).(antlr.TerminalNode).GetText()
			lhs := visitGeneric(v, gv).(*types.PVal)
			if op == "?" {
				b := lhs.GetBool()
				if b {
					rt = visitGeneric(ctx.Expr(1), gv).(*types.PVal)
				} else {
					rt = visitGeneric(ctx.Expr(2), gv).(*types.PVal)
				}
			} else if op == "&&" {
				lhsv := lhs.GetBool()
				if !lhsv {
					rt = types.NewScalarPval(false)
					break
				}
				rhs := visitGeneric(ctx.Expr(1), gv).(*types.PVal)
				rt = types.NewScalarPval(rhs.GetBool())
			} else if op == "||" {
				lhsv := lhs.GetBool()
				if lhsv {
					rt = types.NewScalarPval(true)
					break
				}
				rhs := visitGeneric(ctx.Expr(1), gv).(*types.PVal)
				rt = types.NewScalarPval(rhs.GetBool())
			} else {

				rhs := visitGeneric(ctx.Expr(1), gv).(*types.PVal)
				rt = doOp(ctx, lhs, rhs, op)
			}
		}

	case *parser.SimpleStatementContext:

		switch v := ctx.GetChild(0).(type) {
		case antlr.TerminalNode:
			if v.GetSymbol().GetTokenType() == parser.PinnParserID {
				opString := ctx.GetChildren()[len(ctx.GetChildren())-1].(antlr.TerminalNode).GetText()
				switch opString {
				case "++", "--":
					var result *types.PVal
					v2 := gv.GetPv(ctx.ID().GetText())
					if ctx.Expr() != nil {
						indexE := visitGeneric(ctx.Expr(), gv).(*types.PVal)
						index := types.FixIndex(indexE.GetInterface())
						lhs := types.NewScalarPval(v2.Get(index))
						result = doOp(ctx, lhs, nil, opString)
						v2.Set(index, result.GetInterface())
						break
					}
					result = doOp(ctx, v2, nil, opString)
					gv.PutPv(ctx.ID().GetText(), result)
					break
				default:
					panic(types.ErrCase)
				}

			}
		default:
			visitGeneric(v.(antlr.ParserRuleContext), gv)
		}

	case *parser.ExprListContext:
		if args != nil {
			v := args[0].(*types.PVal)
			rt = false
			for _, child := range ctx.AllExpr() {
				v2 := visitGeneric(child, gv).(*types.PVal)
				if v.Equal(v2) {
					rt = true
					break
				}
			}
		} else {
			rt = make([]*types.PVal, 0, 10)
			for _, child := range ctx.AllExpr() {
				v := visitGeneric(child, gv).(*types.PVal)
				rt = append(rt.([]*types.PVal), v)
			}
		}

	case *parser.CompoundSetContext:
		str := ctx.ID().GetText()
		v := gv.GetPv(str)
		children := ctx.GetChildren()
		op := children[len(children)-3].(antlr.TerminalNode).GetText()
		exprs := ctx.AllExpr()
		if len(exprs) == 2 {
			v2 := visitGeneric(exprs[0], gv).(*types.PVal)
			rhs := visitGeneric(exprs[1], gv).(*types.PVal)
			lhs := types.NewScalarPval(v.Get(types.FixIndex(v2.GetInterface())))
			result := doOp(ctx, lhs, rhs, op)
			v.Set(types.FixIndex(v2.GetInterface()), result.GetInterface())
		} else {
			e := visitGeneric(exprs[0], gv).(*types.PVal)
			result := doOp(ctx, v, e, op)
			v.Set(nil, result.GetInterface())
		}

	case *parser.BlockContext:
		for _, child := range ctx.AllStatement() {
			visitGeneric(child, gv)
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
			e := visitGeneric(ctx.Expr(), gv).(*types.PVal)
			f.M[strID] = e
			break
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
		k := visitGeneric(ctx.Kind(), gv).(types.Kind)
		rt = types.FKind{k, strID}

	case *parser.KindContext:

		strTYPE := ctx.TYPES().GetText()
		t := types.MLitToType[strTYPE]
		if ctx.LSQUARE() == nil {
			rt = types.NewKind(types.GScalar, t, 1)
		} else {
			if ctx.Expr() != nil {
				v := visitGeneric(ctx.Expr(), gv).(*types.PVal)
				x := int(v.GetInt())
				if x == -1 {
					panic(types.ErrRange)
				}
				rt = types.NewKind(types.GArray, t, x)
			} else if ctx.MAP() != nil {
				rt = types.NewKind(types.GMap, t, 0)
			} else if ctx.SLICE() != nil {
				rt = types.NewKind(types.GSlice, t, 0)
			} else if ctx.FILL() != nil {
				rt = types.NewKind(types.GArray, t, -1)
			} else {
				panic(types.ErrCase)
			}
		}

	case *parser.ReturnStatementContext:
		gv.Fc.Path = types.PExiting
		i := ctx.Expr()
		if i != nil {
			gv.Fc.Rt = visitGeneric(i, gv).(*types.PVal)
		}
	case *parser.SimpleSetContext:
		str := ctx.ID().GetText()
		v := gv.GetPv(str)
		exprs := ctx.AllExpr()
		if len(exprs) == 2 {
			v2 := visitGeneric(exprs[0], gv).(*types.PVal)
			e := visitGeneric(exprs[1], gv).(*types.PVal)
			v.Set(types.FixIndex(v2.GetInterface()), e.GetInterface())
		} else {
			e := visitGeneric(exprs[0], gv).(*types.PVal)
			if e.GetKind().GType == types.GScalar {
				v.Set(nil, e.GetInterface())
			} else {
				gv.PutPv(str, e)
			}
		}

	case *parser.IndexExprContext:
		v := gv.GetPv(ctx.ID().GetText())
		if ctx.COLON() != nil {
			lhsv := 0
			if e := ctx.GetFirst(); e != nil {
				lhsv = int(visitGeneric(e, gv).(*types.PVal).GetInt())
			}
			rhsv := v.GetKind().Int
			if e := ctx.GetSecond(); e != nil {
				rhsv = int(visitGeneric(e, gv).(*types.PVal).GetInt())
			}
			rt = v.Slice(lhsv, rhsv)
		} else {
			e := visitGeneric(ctx.Expr(0), gv).(*types.PVal)
			rt = types.NewScalarPval(v.Get(types.FixIndex(e.GetInterface())))
		}
	default:
		panic(types.ErrCase)
	}
	gv.Prc = oldPrc
	return rt

}

func visitFunction(actx antlr.ParserRuleContext, s []*types.PVal, str string, gv *types.GVal) *types.PVal {
	ctx := actx.(*parser.FunctionContext)
	oldPrc := gv.Prc
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
	gv.Prc = oldPrc
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
