Todo
====

* Reorg types to package. X
* Reorg to generic visits.
* Slice strings, reslice slices.

Types
=====

* char. X
* int. int64. no uint, like java. X
* string. arrays of char, immutable. X
* unit. X
* bool. true, false. X
* bigint. signed, good for euclid. X
* bigfloat. X
* "x int" bass ackwards lol X

Lexical
=======

* chars, as defined in types. X
* comments // style X
* identifiers: start with a number? tbd. NO
* WS: space, nl.

integers
--------
* same as go, hex 0xabc, not 0XABC X

char literal:
needed? 'a
string literal:
no need for interpreted. "abc"

slice-only? pass all by ref (doesn't matter for strings) cloning
functions: multi return to start NO
maps... no mucking with cap (same with slice) X

ops

prec: binary, * / % << >> & +(strs) - | + MORE D00D
output bool, == != < <= > >= see golang X
input bool, && || short circuit X

/ % sign rules as java/golang. X
<< allow neg? NO

no label NOPE
no inc/dec DID IT BITCH
no switch to start DID IT

go rox

