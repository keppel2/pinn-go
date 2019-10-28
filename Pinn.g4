grammar Pinn;

TWODOTS : '@' ;
file : ( function | varDecl ';' )+ ;
function
  : 'func' ID LPAREN (fvarDecl (',' fvarDecl)*)?  ')' kind? block ;

testRule
  : name=expr;

block
  : '{' statement* '}' ;

fvarDecl
  : ID kind ;

varDecl
  : ID kind ('=' exprList)?
  | ID CE expr ;

LSQUARE : '[' ;
LPAREN : '(' ;
kind
  : (LSQUARE ( MAP | SLICE | FILL | expr) ']')? TYPES ;

MAP
 : 'map' ;

SLICE
 : 'slice' ;

FILL
 : 'fill' ;

TYPES
  : ('int' | 'bool' | 'unit' | 'string' | 'big' | 'float' | 'char' ) ; 

DOUBLEOP
  : '++' | '--' ;

simpleStatement
  : pset
  | ID (LSQUARE expr ']')? DOUBLEOP ;

indexExpr :
  ID LSQUARE first=expr? (TWODOTS | COLON) second=expr? ']' 
  | ID LSQUARE expr ']' ;

funcExpr
  : 'print' LPAREN exprList ')'
  | 'printB' LPAREN expr ')'
  | 'printH' LPAREN expr ')'
  | 'delete' LPAREN ID ',' expr ')'
  |  'len' LPAREN ID ')'
  | 'strLen' LPAREN expr ')'
  | 'stringValue' LPAREN expr ')' ;

callExpr: 

  ID LPAREN exprList? ')';
parenExpr: LPAREN expr ')';

expr
  : 
  funcExpr
  | indexExpr
  | ('+' | '-' | '!' | '^') expr
  | expr ('+' | '-' | '^' | BINOP) expr
  | expr ('==' | '!=' | '>' | '<' | '>=' | '<=' ) expr
  | expr ('&&' | '||') expr
  | callExpr
  | parenExpr 
  | expr (TWODOTS | COLON) expr
  | ID
  | FLOAT
  | INT
  | BOOL
  | STRING
  | CHAR
  | IOTA

  | expr '?' expr COLON expr;

exprList
  : expr (',' expr)* ;

returnStatement
  : 'return' expr? ';' ;

ifStatement
  : 'if' expr statement ('else' statement)?;

guardStatement
  : 'guard' expr 'else' block;

whStatement
  : 'while' expr block ;

repeatStatement
  : 'repeat' block 'while' expr ;

RANGE : 'range' ;

foStatement
  : 'for' (varDecl | fss=simpleStatement)? ';' expr ';' sss=simpleStatement block
  | 'for' ID COMMA ID '=' RANGE expr block |
    'for' ID '=' RANGE expr block ;
caseStatement
  : 'case' exprList COLON statement* ;

switchStatement
  : 'switch' expr '{' caseStatement* ('default' COLON statement*)? '}' ;

statement
  :  expr ';'
  | varDecl ';' 
  | simpleStatement ';'
  | ifStatement
  | guardStatement
  | whStatement 
  | repeatStatement ';'
  | switchStatement
  | returnStatement
  | foStatement
  | block
  | 'break' ';'
  | 'continue' ';'
  | 'fallthrough' ';'
  | 'debug' LPAREN ')' ';'
  | ';' ;
pset
  : ID (LSQUARE expr ']')? '=' expr #simpleSet
  | ID (LSQUARE index=expr ']')? ('+' | '-' | '^' | BINOP) '=' rhs=expr #compoundSet ;

COMMA : ',' ;
COLON : ':' ;
CE : ':=' ;
IOTA : 'iota' ;
BINOP : ('*' | '/' | '%' | '&' | '|' | '<<' | '>>' ) ;
BOOL : 'true' | 'false' ;
ID : [a-zA-Z_]([a-zA-Z_0-9])* ;
CHAR : '\''[a-zA-Z_0-9]'\'' ;
INT : '0'
  | [1-9] ('_'? DECIMAL_DIGITS)?
  | '0x' '_'? HEX_DIGIT ('_'? HEX_DIGIT)*
  | '0b' '_'? BINARY_DIGIT ('_'? BINARY_DIGIT)*
  | '0o' '_'? OCTAL_DIGIT ('_'? OCTAL_DIGIT)* ;

FLOAT : DECIMAL_DIGITS '.' DECIMAL_DIGITS? DECIMAL_EXPONENT?
  | DECIMAL_DIGITS DECIMAL_EXPONENT
  | '.' DECIMAL_DIGITS DECIMAL_EXPONENT?
  | '0x' HEX_MANTISSA HEX_EXPONENT ;

WS : ([ \t\n]+ | '//' ~('\n')* '\n' | '/*' .*? '*/' )-> skip ;
STRING : '"' ~('"')* '"' ;

fragment DECIMAL_DIGIT : [0-9] ;
fragment DECIMAL_DIGITS : DECIMAL_DIGIT ('_'? DECIMAL_DIGIT)* ;
fragment DECIMAL_EXPONENT : 'e' [+-]? DECIMAL_DIGITS ;
fragment HEX_DIGIT : [0-9a-f] ;
fragment HEX_DIGITS: HEX_DIGIT ('_'? HEX_DIGIT)* ;
fragment HEX_MANTISSA : '_'? HEX_DIGITS '.' HEX_DIGITS?
  | '_'? HEX_DIGITS
  | '.' HEX_DIGITS ;
fragment HEX_EXPONENT : 'p' [+-]? DECIMAL_DIGITS 
  | 'h' [+-]? HEX_DIGITS ;
fragment OCTAL_DIGIT : [0-7] ;
fragment BINARY_DIGIT : [01] ;
