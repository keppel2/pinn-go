package main

import "github.com/antlr/antlr4/runtime/Go/antlr"
import "./pparser"
import "./types"

import "fmt"
import "strconv"

import "strings"
import "io/ioutil"

//import "io"
import "os"

import "math/big"

import "flag"

//import "encoding/json"

var bListTokens = flag.Bool("token", false, "desc")
var sFile = flag.String("f", "a.pinn", "file")
var bDebug = flag.Bool("d", false, "debug")
var bOne = flag.Bool("one", false, "one-off")

var _ = strconv.Atoi
var _ = strings.Replace

func doOp(ctx antlr.ParserRuleContext, lhs, rhs *types.PVal, op string) *types.PVal {
	var rt *types.PVal
	switch op {
	case "+", "-", "*", "/", "%", "&", "|", "^", "<", ">", "<=", ">=", "<<", ">>":
		switch lhs.GetKind().VType {
		case types.TString:
			lhsv := lhs.GetString()
			rhsv := rhs.GetString()
			switch op {
			case "+":
				rt = types.NewScalarPval(lhsv + rhsv)
			case "<":
				rt = types.NewScalarPval(lhsv < rhsv)
			case "<=":
				rt = types.NewScalarPval(lhsv <= rhsv)
			case ">":
				rt = types.NewScalarPval(lhsv > rhsv)
			case ">=":
				rt = types.NewScalarPval(lhsv >= rhsv)
			default:
				panic(types.ErrCase)
			}
		case types.TBig:
			lhsv := lhs.GetBig()
			if op == "<<" {
				return types.NewScalarPval(new(big.Int).Lsh(lhsv, types.ToUint(rhs.GetInterface())))
			}
			if op == ">>" {
				return types.NewScalarPval(new(big.Int).Rsh(lhsv, types.ToUint(rhs.GetInterface())))
			}
			rhsv := types.TBig.Convert(rhs.GetInterface()).(*big.Int)
			switch op {
			case "+":
				rt = types.NewScalarPval(new(big.Int).Add(lhsv, rhsv))
			case "-":
				rt = types.NewScalarPval(new(big.Int).Sub(lhsv, rhsv))
			case "*":
				rt = types.NewScalarPval(new(big.Int).Mul(lhsv, rhsv))
			case "/":
				rt = types.NewScalarPval(new(big.Int).Div(lhsv, rhsv))
			case "%":
				rt = types.NewScalarPval(new(big.Int).Mod(lhsv, rhsv))
			case "&":
				rt = types.NewScalarPval(new(big.Int).And(lhsv, rhsv))
			case "|":
				rt = types.NewScalarPval(new(big.Int).Or(lhsv, rhsv))
			case "^":
				rt = types.NewScalarPval(new(big.Int).Xor(lhsv, rhsv))
			case "<":
				rt = types.NewScalarPval(lhsv.Cmp(rhsv) == -1)
			case "<=":
				rt = types.NewScalarPval(lhsv.Cmp(rhsv) != 1)
			case ">":
				rt = types.NewScalarPval(lhsv.Cmp(rhsv) == 1)
			case ">=":
				rt = types.NewScalarPval(lhsv.Cmp(rhsv) != -1)
			default:
				panic(types.ErrCase)
			}

		case types.TBigf:
			lhsv := lhs.GetBigf()
			rhsv := types.TBigf.Convert(rhs.GetInterface()).(*big.Float)
			switch op {
			case "+":
				rt = types.NewScalarPval(new(big.Float).Add(lhsv, rhsv))
			case "-":
				rt = types.NewScalarPval(new(big.Float).Sub(lhsv, rhsv))
			case "*":
				rt = types.NewScalarPval(new(big.Float).Mul(lhsv, rhsv))
			case "/":
				rt = types.NewScalarPval(new(big.Float).Quo(lhsv, rhsv))
			case "%":
				panic(types.ErrWrongType)
			case "<":
				rt = types.NewScalarPval(lhsv.Cmp(rhsv) == -1)
			case "<=":
				rt = types.NewScalarPval(lhsv.Cmp(rhsv) != 1)
			case ">":
				rt = types.NewScalarPval(lhsv.Cmp(rhsv) == 1)
			case ">=":
				rt = types.NewScalarPval(lhsv.Cmp(rhsv) != -1)
			default:
				panic(types.ErrCase)
			}
		case types.TChar:
			lhsv := lhs.GetChar()
			rhsv := rhs.GetChar()
			switch op {
			case "+":
				rt = types.NewScalarPval(lhsv + rhsv)
			case "-":
				rt = types.NewScalarPval(lhsv - rhsv)
			case "*":
				rt = types.NewScalarPval(lhsv * rhsv)
			case "/":
				rt = types.NewScalarPval(lhsv / rhsv)
			case "%":
				rt = types.NewScalarPval(lhsv % rhsv)
			case "&":
				rt = types.NewScalarPval(lhsv & rhsv)
			case "|":
				rt = types.NewScalarPval(lhsv | rhsv)
			case "^":
				rt = types.NewScalarPval(lhsv ^ rhsv)
			case "<<":
				rt = types.NewScalarPval(lhsv << rhsv)
			case ">>":
				rt = types.NewScalarPval(lhsv >> rhsv)
			case "<":
				rt = types.NewScalarPval(lhsv < rhsv)
			case "<=":
				rt = types.NewScalarPval(lhsv <= rhsv)
			case ">":
				rt = types.NewScalarPval(lhsv > rhsv)
			case ">=":
				rt = types.NewScalarPval(lhsv >= rhsv)
			default:
				panic(types.ErrCase)
			}

		case types.TInt:
			lhsv := lhs.GetInt()
			rhsv := rhs.GetInt()
			switch op {
			case "+":
				rt = types.NewScalarPval(lhsv + rhsv)
			case "-":
				rt = types.NewScalarPval(lhsv - rhsv)
			case "*":
				rt = types.NewScalarPval(lhsv * rhsv)
			case "/":
				rt = types.NewScalarPval(lhsv / rhsv)
			case "%":
				rt = types.NewScalarPval(lhsv % rhsv)
			case "&":
				rt = types.NewScalarPval(lhsv & rhsv)
			case "|":
				rt = types.NewScalarPval(lhsv | rhsv)
			case "^":
				rt = types.NewScalarPval(lhsv ^ rhsv)
			case "<<":
				rt = types.NewScalarPval(lhsv << rhsv)
			case ">>":
				rt = types.NewScalarPval(lhsv >> rhsv)
			case "<":
				rt = types.NewScalarPval(lhsv < rhsv)
			case "<=":
				rt = types.NewScalarPval(lhsv <= rhsv)
			case ">":
				rt = types.NewScalarPval(lhsv > rhsv)
			case ">=":
				rt = types.NewScalarPval(lhsv >= rhsv)
			default:
				panic(types.ErrCase)
			}
		default:
			panic(types.ErrCase)
		}
	case "==":
		rt = types.NewScalarPval(lhs.Equal(rhs))
	case "!=":
		rt = types.NewScalarPval(!lhs.Equal(rhs))
	case "--":
		switch lhs.GetKind().VType {
		case types.TBig:
			rt = types.NewScalarPval(new(big.Int).Sub(lhs.GetBig(), big.NewInt(1)))
		case types.TInt:
			rt = types.NewScalarPval(lhs.GetInt() - 1)
		case types.TBigf:
			rt = types.NewScalarPval(new(big.Float).Sub(lhs.GetBigf(), new(big.Float).SetInt64(1)))
		default:
			panic(types.ErrCase)
		}
	case "++":
		switch lhs.GetKind().VType {
		case types.TBig:
			rt = types.NewScalarPval(new(big.Int).Add(lhs.GetBig(), big.NewInt(1)))
		case types.TInt:
			rt = types.NewScalarPval(lhs.GetInt() + 1)
		case types.TBigf:
			rt = types.NewScalarPval(new(big.Float).Add(lhs.GetBigf(), new(big.Float).SetInt64(1)))
		default:
			panic(types.ErrCase)
		}
	default:
		panic(types.ErrCase)
	}
	return rt

}

func floatToDec(s string) string {
	strs := strings.Split(s, "h")
	if len(strs) == 1 {
		return s
	}
	hexStr := strs[1]
	i, _ := strconv.ParseInt(hexStr, 16, 64)
	decStr := fmt.Sprint(i)
	return strs[0] + "p" + decStr
}

func floatToHex(s string) string {
	strs := strings.Split(s, "p")
	decStr := strs[1]
	i, _ := strconv.ParseInt(decStr, 10, 64)
	hexStr := fmt.Sprintf("%x", i)
	return strs[0] + "h" + hexStr
}

func symTokens(
	ctx interface {
		antlr.ParserRuleContext
		GetParser() antlr.Parser
	}) ([]antlr.Token, []antlr.Token) {
	rtTok := make([]antlr.Token, 0)
	rtLit := make([]antlr.Token, 0)
	children := ctx.GetChildren()
	for _, child := range children {
		if v, ok := child.(antlr.TerminalNode); ok {
			tok := v.GetSymbol()
			if ctx.GetParser().GetSymbolicNames()[tok.GetTokenType()] != "" {
				rtTok = append(rtTok, tok)
			} else {
				rtLit = append(rtLit, tok)
			}
		}
	}
	return rtTok, rtLit
}

func nonTerminals(ctx antlr.ParserRuleContext) []antlr.ParserRuleContext {
	rt := make([]antlr.ParserRuleContext, 0)
	children := ctx.GetChildren()
	for _, child := range children {
		if v, ok := child.(antlr.ParserRuleContext); ok {
			rt = append(rt, v)
		}
	}
	return rt
}

func valAndType(i interface{}) {
	fmt.Fprintln(os.Stderr, fmt.Sprintf("i:%#v", i), fmt.Sprintf("t:%T", i))
}

func dbg(gv *types.GVal) {
	str := fmt.Sprintln("l:", gv.Prc.GetStart().GetLine(), "c:", gv.Prc.GetStart().GetColumn())
	str += fmt.Sprintln("t:", gv.Prc.GetText(), "g:", gv.Gfc.M, "l:", gv.Fc.M)
	str += fmt.Sprintln(fmt.Sprintf("ld:%#v", gv.Fc.M), fmt.Sprintf("gd:%#v", gv.Gfc.M))
	for i := len(gv.Stack) - 1; i >= 0; i-- {
		str += fmt.Sprintln("s", i, ":", gv.Stack[i].GetStart().GetLine())
	}
	fmt.Fprint(os.Stderr, str)
}

func fileToString(s string) string {
	bytes, err := ioutil.ReadFile(s)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func stringToParser(s string) *parser.PinnParser {
	input := antlr.NewInputStream(s)
	lexer := parser.NewPinnLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	return parser.NewPinnParser(stream)
}

func parse(fstr string) *parser.FileContext {
	p := stringToParser(fstr)
	p.SetErrorHandler(antlr.NewBailErrorStrategy())
	defer func() { recover() }()
	tree := p.File()
	return tree.(*parser.FileContext)
}

func parseForErrors(fstr string) {
	stringToParser(fstr).File()
}

func reset() {
}

func exe(fstr string, b bool) {
	pfc := parse(fstr)
	if pfc == nil {
		parseForErrors(fstr)
		return
	}
	visitFile(pfc, b)
}

func oneF() bool { println("oneF called"); return true }
func oneOff() {
	switch {
	case true:
		fallthrough
	case true, oneF():
		println("swiii")
	}
}

func main() {
	flag.Parse()
	if *bOne {
		oneOff()
		return
	}
	if *bListTokens {
		/*
			s := make([]string, 0, 800)
			stream := stream(*sFile)
			for _, token := range stream.GetAllTokens() {
				s = append(s, token.GetText())
			}
			fmt.Println(s)
		*/
		panic(types.ErrWrongStatement)
		return
	}

	fstr := fileToString(*sFile)
	exe(fstr, true)
}

type stringPrinter string

type stringOnly string

func (so *stringOnly) print(s string) {
	str := string(*so)
	str = string(str) + s
	*so = stringOnly(str)
}

type gPrinter interface {
	print(string)
}

var globalPrinter gPrinter = new(stringPrinter)

func (sp *stringPrinter) print(s string) {
	fmt.Print(s)
	str := string(*sp)
	str = string(str) + s
	*sp = stringPrinter(str)
}
