grammar Turnkey;

// Whitespace
WS: [ \r\n\t]+ -> skip;
COMMA: ',';

// Keywords
FUNC: 'func';
TURN: 'turn';

// Operators
MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';

ASSIGN: '=';
LPAREN: '(';
RPAREN: ')';
LBRACE: '{';
RBRACE: '}';

// Literals
INT: [1-9][0-9]*;
FLOAT: [1-9][0-9]* '.' [0-9]+;
BOOL: 'true' | 'false';

IDENT: [_]*[A-Za-z][A-Za-z0-9_]*;

// Rules
start : statement EOF;

// Statements
statement
    : expression
    | func_def
;

block
    : LBRACE expression* RBRACE
;

// Functions
func_def
    : FUNC IDENT LPAREN parameters RPAREN block
;

parameters
    : (IDENT COMMA)* IDENT?
;

// Expressions
expression
    : int_expression
    | float_expression
    | bool_expression
    | ident_expression
    | call_expression
;

int_expression
    : int_expression op=('*'|'/') int_expression
    | int_expression op=('+'|'-') int_expression
    | INT
;

float_expression
    : float_expression op=('*'|'/') float_expression
    | float_expression op=('+'|'-') float_expression
    | FLOAT
;

bool_expression
    : BOOL
;

ident_expression
    : IDENT
;

call_expression // Call a function
    : IDENT LPAREN parameters RPAREN
;