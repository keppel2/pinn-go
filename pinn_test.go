package main

import "testing"
import "io/ioutil"

//import "strings"

/*
func TestParse(t *testing.T) {
	files, _ := ioutil.ReadDir(".")
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".pinn") {
			defer func() {
				if e := recover(); e != nil {
					t.Fatal(file.Name())
				}
			}()
			parse(file.Name(), true)
		}
	}
}
*/

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
func TestFunky(t *testing.T) {
	testFile(t, "funky")
}

func TestQs(t *testing.T) {
	testFile(t, "qs")
}
