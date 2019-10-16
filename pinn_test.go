package main

import "testing"
import "io/ioutil"

func testFile(t *testing.T, prefix string) {
	println(prefix)
	globalPrinter = new(stringOnly)
	exe(fileToString(prefix+".pinn"), true)
	str := string(*globalPrinter.(*stringOnly))
	/*
		if testing.Verbose() {
			str = string(*globalPrinter.(*stringPrinter))
		}
	*/

	bytes, _ := ioutil.ReadFile(prefix + ".gold")
	strGold := string(bytes)
	if str != strGold {
		t.Fatal(prefix)
	}
}

func testNeg(t *testing.T, s string) {
	defer func() { recover() }()
	exe(s, false)
	t.Fatal(s)
}

func TestNeg(t *testing.T) {
	testNeg(t,
		`
func main()
{
	i = 5;
}
`)
	testNeg(t,
		`
func nomain()
{
}
`)
	testNeg(t,
		`
func v() {}
func v() {}
func main()
{
}
`)
	testNeg(t,
		`
func main()
{
	guard false else {
	}
}
`)

	testNeg(t,
		`
func v() {}
func main()
{
	str := "string";
	i int;
	i = str;
}
`)

}

func TestBig(t *testing.T) {
	testFile(t, "big")
}

func TestExpr(t *testing.T) {
	testFile(t, "expr")
}

func TestControl(t *testing.T) {
	testFile(t, "control")
}
func TestData(t *testing.T) {
	testFile(t, "data")
}
func TestFunc(t *testing.T) {
	testFile(t, "func")
}

func TestQs(t *testing.T) {
	testFile(t, "qs")
}
