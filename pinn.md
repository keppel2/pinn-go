Pinn
====

## Overview

Pinn is a computer language that is heavily influenced by Go. It is statically typed and imperative. Some aims include:

* Syntax taken from a variety of languages. Unopiniated about which the user should use.
* Markedly reduced numeric types.

## Specification

The grammar is clean of implementation language and is written in ANTLR. It happens to be interpreted in the Go language for now.

Running
=======

Save the following to `a.pinn`:

> func main ( ) {
> 	print ("Hello, world.");
> }

Then compile the ANTLR code and run `go run -f a.pinn`. You should see:

> Hello, world.

