lexer grammar TurnkeyLexer;

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

// Strings...
OPEN_STRING: '"' -> pushMode(IN_STRING);

mode IN_STRING;

STRING: ~[\\"]+ ;
CLOSE_STRING: '"' -> popMode;