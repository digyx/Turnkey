parser grammar TurnkeyParser;

options { tokenVocab=TurnkeyLexer; }

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
    | string_expression
    | ident_expression
    | call_expression
    | turn_expression
;

// Literals
int_expression
    : LPAREN int_expression RPAREN
    | int_expression op=(MUL|DIV) int_expression
    | int_expression op=(ADD|SUB) int_expression
    | INT
;

float_expression
    : LPAREN float_expression RPAREN
    | float_expression op=('*'|'/') float_expression
    | float_expression op=('+'|'-') float_expression
    | FLOAT
;

bool_expression
    : BANG bool_expression
    | bool_expression op=(EQ|NOT_EQ) bool_expression
    | BOOL
;

string_expression
    : string_expression op=ADD string_expression
    | OPEN_STRING STRING CLOSE_STRING
;

ident_expression
    : IDENT
;

call_expression // Call a function
    : IDENT LPAREN parameters RPAREN
;

turn_expression
    : TURN call_expression
;