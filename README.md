# Pinn is a computer language that lets you be correct.

Pinn is a statically typed, imperative language for computers.

* Syntax taken mostly from Go.
* Carefully chosen ideas from Swift.

## Specification

## Lexical

Most lexical elements are borrowed from Go. Later elements include support for `_` in numeric literals and binary floating point.

## Element Types

* `int`. Always a signed 64-bit.
* `bool`. Standard, `true` or `false`.
* `string`. Immutable.
* `big`. Big integer.
* `float`. Big float. See Go.
* `char`. Character. Implemented as UTF-32.

## Group Types

* Scalar is a single element.
* `map`. Key is always a string. Like Go, a missing string returns the zero value.
* `slice`. Pointer to a region of an array or slice. Shares memory in Go and copies in Swift.
* `array`. Arrays are passed by value.

## Expressions

* Almost all taken from Go, so much like c/Java. Conditional expression was put back in.

## Compilation unit
* `( <function> | <statement>)+ EOF`

## block
* `"{" { <statement> } "}"`

## statement

* `while <expr> <block>`
  * Evaluate `expr`. If true, execute `block` and repeat this line. If false, go on.
* `repeat <block> while <expr>`
  * Execute `block`. Evaluate `expr`. If true, repeat this line.
* `return [<expr>] ;`
  * Return from function. The `expr` must match the return type, or empty if there is none. If global, there is no return type.
* `if <expr> <statement> [else <statement>]`
  * Evaluate `expr`. If true, exectue first `statement`. If false, either move on or execute second `statement`.
* `guard <expr> else <block>`
  * Evaluate `expr`. If false, execute `block`. The block must relinquish control, with a `return`, `break`, or `continue`.

## Variable declaration
* `<id> <kind> [ = <expr_list> ]`
  * Declare `id` of `kind` type. Optionally initialize it.
* `id := <expr>`
  * Short declaration.


The grammar is clean of implementation language and is written in ANTLR. It has implementations in Go and Swift. The Swift implementation is more recent and at https://github.com/keppel2/pinnSwift.

# Running Go version


Get ANTLR from https://www.antlr.org/download.html.

Save the following to `hello.pinn`:

> func main ( ) {
> 	print ("Hello, world.");
> }

Then run the ANTLR code and run `go run -f hello.pinn`. You should see:

> Hello, world.

Running ANTLR
-------------

* `java -jar <path_to_antlr_jar> -Dlanguage=Go -o pparser Pinn.g4`

# Elements
## Notation
```
|  alternation
()  grouping
[]  option (0 or 1 times)
{}  repetition (0 to n times)
literal  type as specified
<production>  rule specified elsewhere
"{" "[" ...  notation elements as literals
```

