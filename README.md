Pinn
====

## Overview

Pinn is a computer language that is heavily influenced by Go. It is statically typed and imperative. Some aims include:

* Syntax taken from a variety of languages. Unopiniated about which the user should use.
* Markedly reduced numeric types.
* Built-in numeric big and float types.


## Specification

The grammar is clean of implementation language and is written in ANTLR. It happens to be interpreted in the Go language for now.

Running
=======

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

## Lexical

Most lexical elements are borrowed from Go.

## Element Types

* `int`. Always a signed 64-bit.
* `bool`. Standard, `true` or `false`.
* `unit`. Can only be `iota`. As a special case, getting an element from a map of unit will return `true` if the key is defined (to `iota`), and `false` otherwise.
* `string`. Immutable. Indexed into characters.
* `big`. Big integer.
* `float`. Big float. See Go.
* `char`. Character. Implemented as UTF-32.

## Group Types

* Scalar is a single element.
* `map`. Key is always a string.
* `slice`. Pointer to a region of an array or slice. See Go.
* `array`. Arrays are passed by value.

## Expressions

* Almost all taken from Go, so much like c/Java. Conditional expression was put back in.

## Statements

Taken from Go, with some variations.
