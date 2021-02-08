grammar Turnkey;

// Whitespace
WS: [ \r\n\t]+ -> skip;

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

// Literals
INT: [1-9][0-9]*;
FLOAT: [1-9][0-9]* '.' [0-9]+;
BOOL: 'true' | 'false';
STRING: '"' [A-Za-z0-9]+ '"';

IDENT: [_]*[a-z][A-Za-z0-9_]*;

// Rules
start : expression EOF;

expression
    : int_expression
    | float_expression
    | string_expression
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

string_expression
    : string_expression op='+' string_expression
    | STRING
;