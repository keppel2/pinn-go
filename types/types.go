package types

import "github.com/antlr/antlr4/runtime/Go/antlr"
import "../pparser"

import "fmt"

import "errors"
import "math/big"
var (
	ErrParamLength    = errors.New("Parameter length mismatch.")
	ErrWrongStatement = errors.New("Wrong statement.")
	ErrWrongType      = errors.New("Wrong type.")
	ErrRange          = errors.New("Out of range.")
	ErrCase           = errors.New("Case unimplemented.")
	ErrReDeclare      = errors.New("Redeclared.")
	ErrUnDeclare      = errors.New("Undeclared.")
)

type GType int

const (
	gInvalid GType = iota
	GScalar
	GArray
	GSlice
	GMap
)


type VType int

const (
	tInvalid VType = iota
	TInt
	TBool
	TUnit
	TString
	TBig
	TBigf
	TChar
)

func (v VType) newType() interface{} {
	var rt interface{}
	switch v {
	case TInt:
		rt = int64(0)
	case TBool:
		rt = false
	case TUnit:
		rt = nil
	case TString:
		rt = ""
	case TBig:
		rt = new(big.Int)
	case TBigf:
		rt = new(big.Float)
	case TChar:
		rt = rune(0)
	default:
		panic(ErrCase)
	}
	return rt

}

func (v VType) sConvert(i interface{}) string {
	switch v {
	case TUnit:
		return "iota"
	case TChar:
		return string(i.(rune))
	default:
		panic(ErrCase)

	}
}

func (v VType) Convert(i interface{}) interface{} {
	rt := i
	v2 := intToVtype(rt)
	switch v {
	case TBig:
		switch v2 {
		case TInt:
			rt = big.NewInt(rt.(int64))
		}
	case TBigf:
		switch v2 {
		case TInt:
			rt = new(big.Float).SetInt64(rt.(int64))
		case TBig:
			rt = new(big.Float).SetInt(rt.(*big.Int))
		}
	}
	if v != intToVtype(rt) {
		panic(ErrCase)
	}
	return rt
}

func ToUint(i interface{}) uint {
	switch val := i.(type) {
	case *big.Int:
		return uint(val.Uint64())
	case int64:
		return uint(val)
	default:
		panic(ErrWrongType)
	}
}

func copyBig(b *big.Int) *big.Int {
	return new(big.Int).Set(b)
}

type GVal struct {
	Fc *Fc
	Fkmap map[string]Fheader
	Gfc   *Fc
	Prc   antlr.ParserRuleContext
	Stack []antlr.ParserRuleContext
}

type Fheader struct {
	*parser.FunctionContext
	Kind
	S []FKind
}

func (g GVal) IsPathNormal() bool { return g.Fc.Path == PNormal }

func (g GVal) PathEndBlock() bool {
	switch g.Fc.Path {
	case PContinue, PNormal:
		g.Fc.Path = PNormal
		return false
	case PBreak:
		g.Fc.Path = PNormal
		fallthrough
	case PExiting:
		return true
	default: panic(ErrCase)
	}
}

func (g *GVal) PushFrame() {
	g.Stack = append(g.Stack, g.Prc)
}

func (g GVal) PutPv(s string, pv *PVal) {
	for _, fc := range []*Fc{g.Fc, g.Gfc} {
		if fc == nil {
			continue
		}
		if oldPv, ok := fc.M[s]; ok {
			pv.GetKind().EqualsError(oldPv.GetKind())
			fc.M[s] = pv
			return
		}
	}
	panic(ErrWrongType)
}

func (g GVal) GetPv(s string) *PVal {
	for _, v := range []*Fc{g.Fc, g.Gfc} {
		if v == nil {
			continue
		}
		if v2, ok := v.M[s]; ok {
			return v2
		}
	}
	panic(ErrUnDeclare)
}

type PVal struct {
	v VType
	g GType
	a []interface{}
	m map[string]interface{}
	i interface{}
	fmt.Stringer
	fmt.GoStringer
}

func (p PVal) GoString() string {
	str := ""
	switch p.g {
	case GScalar:
		str = fmt.Sprintf("v:%v. g:%v. i:%v. &i:%v", p.v, p.g, p.i, &p.i)
	case GArray, GSlice:
		str = fmt.Sprintf("v:%v. g:%v. a:%v", p.v, p.g, p.a)
	case GMap:
		str = fmt.Sprintf("v:%v. g:%v. m:%v", p.v, p.g, p.m)

	default:
		panic(ErrCase)
	}
	return str
}

func (p PVal) Slice(lhs int, rhs int) *PVal {
	rt := NewPval(GSlice, p.v, rhs-lhs)
	rt.a = p.a[lhs:rhs]
	return rt
}

func (p *PVal) Invalid() bool {
	return p == nil 
}

func (p PVal) GetKeys() []string {
	rt := make([]string, 0, len(p.m))
	for k, _ := range p.m {
		rt = append(rt, k)
	}
	return rt
}

func (p PVal) GetKind() Kind {
	switch p.g {
	case GArray, GSlice:
		return NewKind(p.g, p.v, len(p.a))
	case GMap:
		return NewKind(p.g, p.v, len(p.m))
	case GScalar:
		return NewKind(p.g, p.v, 1)
	default:
		panic(ErrCase)
	}
}

func (p *PVal) DeleteKey(k string) {
	if p.g != GMap {
		panic(ErrWrongType)
	}
	delete(p.m, k)
}
func (p *PVal) Set(k interface{}, v interface{}) {
	switch p.g {
	case GScalar:
		if k != nil {
			panic(ErrWrongType)
		}
		p.i = p.GetKind().VType.Convert(v)
	case GSlice:
		if k.(int) == len(p.a) {
			p.a = append(p.a, p.GetKind().VType.Convert(v))
			return
		}
		fallthrough
	case GArray:
		p.a[k.(int)] = p.GetKind().VType.Convert(v)
	case GMap:
		p.m[k.(string)] = p.GetKind().VType.Convert(v)
	default:
		panic(ErrCase)
	}
}

func (p PVal) Get(k interface{}) interface{} {
	switch p.g {
	case GScalar:
		return p.i
	case GArray, GSlice:
		return p.a[k.(int)]
	case GMap:
		if p.GetKind().VType == TUnit {
			_, ok := p.m[k.(string)]
			return ok
		}
		return p.m[k.(string)]
	default:
		panic(ErrCase)
	}
}

func typeEqual(i1 interface{}, i2 interface{}, v VType) bool {
	switch v {
	case TInt:
		return i1.(int64) == i2.(int64)
	case TBool:
		return i1.(bool) == i2.(bool)
	case TString:
		return i1.(string) == i2.(string)
	case TBig:
		return i1.(*big.Int).Cmp(i2.(*big.Int)) == 0
	case TBigf:
		return i1.(*big.Float).Cmp(i2.(*big.Float)) == 0
	case TChar:
		return i1.(rune) == i2.(rune)
	case TUnit:
		return true
	default:
		panic(ErrCase)
	}
}

func copyType(in interface{}) interface{} {
	v := intToVtype(in)
	var rt interface{}
	switch v {
	case TInt:
		rt = in.(int64)
	case TBool:
		rt = in.(bool)
	case TUnit:
		rt = nil
	case TString:
		rt = in.(string)
	case TBig:
		rt = new(big.Int).Set(in.(*big.Int))
	case TBigf:
		rt = new(big.Float).Set(in.(*big.Float))
	case TChar:
		rt = in.(rune)
	default:
		panic(ErrCase)
	}
	return rt
}

func (p PVal) stringConvert() PVal {
	rt := NewPval(p.g, TString, p.GetKind().Int)
	switch p.g {
	case GScalar:
		rt.i = p.v.sConvert(p.i)
	case GArray, GSlice:
		for k, v := range p.a {
			rt.a[k] = p.v.sConvert(v)
		}
	case GMap:
		for k, v := range p.m {
			rt.m[k] = p.v.sConvert(v)
		}
	default:
		panic(ErrCase)
	}
	return *rt
}
func (p PVal) Clone() *PVal {
	rt := NewPval(p.g, p.v, p.GetKind().Int)
	switch p.g {
	case GScalar:
		rt.i = copyType(p.i)
	case GArray:
		for k, v := range p.a {
			rt.a[k] = copyType(v)
		}
	case GMap:
	default:
		panic(ErrCase)
	}
	return rt
}

func (p PVal) Equal(p2 *PVal) bool {
	switch p.g {
	case GMap, GSlice:
		panic(ErrWrongType)
	}
	p.GetKind().EqualsError(p2.GetKind())
	switch p.GetKind().GType {
	case GScalar:
		return typeEqual(p.i, p2.i, p.GetKind().VType)
	case GArray:
		for x := 0; x < p.GetKind().Int; x++ {
			if !typeEqual(p.a[x], p2.a[x], p.GetKind().VType) {
				return false
			}
		}
		return true
	default:
		panic(ErrCase)
	}
}

func (p PVal) String() string {
	if p.GetKind().VType == TChar || p.GetKind().VType == TUnit {
		p = p.stringConvert()
	}
	switch p.g {
	case GScalar:
		return fmt.Sprint(p.i)
	case GArray, GSlice:
		return fmt.Sprint(p.a)
	case GMap:
		return fmt.Sprint(p.m)
	default:
		panic(ErrCase)
	}
}

func (p PVal) GetInt() int64 {
	return p.GetInterface().(int64)
}
func (p PVal) GetChar() rune {
	return p.GetInterface().(rune)
}
func (p PVal) GetBool() bool {
	return p.GetInterface().(bool)
}

func (p PVal) GetString() string {
	return p.GetInterface().(string)
}

func (p PVal) GetBig() *big.Int {
	return p.GetInterface().(*big.Int)
}

func (p PVal) GetBigf() *big.Float {
	return p.GetInterface().(*big.Float)
}
func (p PVal) GetInterface() interface{} {
	if p.g != GScalar {
		panic(ErrWrongType)
	}
	return p.Get(nil)
}

func intToVtype(input interface{}) VType {
	if input == nil {
		return TUnit
	}
	switch input.(type) {
	case int64:
		return TInt
	case bool:
		return TBool
	case string:
		return TString
	case *big.Int:
		return TBig
	case *big.Float:
		return TBigf
	case rune:
		return TChar
	default:
		panic(ErrCase)
	}
}

func NewScalarPval(input interface{}) *PVal {
	rt := NewPval(GScalar, intToVtype(input), 1)
	rt.i = input
	return rt
}

func NewPval(g GType, v VType, size int) *PVal {
	rt := &PVal{g: g, v: v}
	switch g {
	case GScalar:
		switch v {
		case TInt:
			rt.i = int64(0)
		case TBool:
			rt.i = false
		case TUnit:
			rt.i = nil
		case TString:
			rt.i = ""
		case TBig:
			rt.i = new(big.Int)
		case TBigf:
			rt.i = new(big.Float)
		case TChar:
			rt.i = rune(0)
		default:
			panic(ErrCase)
		}
	case GArray, GSlice:
		rt.a = make([]interface{}, size)
		for k, _ := range rt.a {
			rt.a[k] = v.newType()
		}
	case GMap:
		rt.m = make(map[string]interface{}, 24)
	default:
		panic(ErrCase)
	}
	return rt
}

var MLitToType = map[string]VType{
	"int":    TInt,
	"bool":   TBool,
	"unit":   TUnit,
	"string": TString,
	"big":    TBig,
	"float":  TBigf,
	"char":   TChar}

type FKind struct {
	Kind
	S string
}

type Kind struct {
	VType
	GType
	Int int
}

func (k Kind) EqualsError(k2 Kind) {
	switch k.GType {
	case GArray, GScalar:
		if k != k2 {
			panic(ErrWrongType)
		}
	case GSlice:
		if k2.GType != GSlice || k.VType != k2.VType {
			panic(ErrWrongType)
		}
	case GMap:
		if k2.GType != GMap || k.VType != k2.VType {
			panic(ErrWrongType)
		}
	default:
		panic(ErrCase)
	}
}

func NewKind(g GType, v VType, x int) Kind {
	rt := Kind{GType: g, VType: v, Int: x}
	return rt
}

func (k Kind) Iv() *PVal {
	return NewPval(k.GType, k.VType, k.Int)
}

type path int

const (
	PNormal path = iota
	PExiting
	PBreak
	PContinue
	PFallthrough
)

func (p path) PathToNormal() path {
	switch p {
	case PContinue, PBreak:
		return PNormal
	case PExiting, PNormal:
		return p
	default:
		panic(ErrCase)
	}
}

type Fc struct {
	M         map[string]*PVal
	Path path
	Rt       *PVal
}

func NewFc() *Fc {
	rt := Fc{
		M: make(map[string]*PVal, 10)}
	return &rt
}

func (f Fc) Clone() *Fc {
	rt := f
	rt.M = make(map[string]*PVal, 10)
	for k, v := range f.M {
		if v.GetKind().GType != GSlice && v.GetKind().GType != GMap {
			rt.M[k] = v.Clone()
		} else {
			rt.M[k] = v
		}
	}
	return &rt
}

func FixIndex(i interface{}) interface{} {
	switch rt := i.(type) {
	case int64:
		return int(rt)
	case string:
		return rt
	default:
		panic(ErrCase)
	}
}



