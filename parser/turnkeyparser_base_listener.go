// Code generated from TurnkeyParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // TurnkeyParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTurnkeyParserListener is a complete listener for a parse tree produced by TurnkeyParser.
type BaseTurnkeyParserListener struct{}

var _ TurnkeyParserListener = &BaseTurnkeyParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTurnkeyParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTurnkeyParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTurnkeyParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTurnkeyParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseTurnkeyParserListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseTurnkeyParserListener) ExitStart(ctx *StartContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseTurnkeyParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseTurnkeyParserListener) ExitStatement(ctx *StatementContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseTurnkeyParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseTurnkeyParserListener) ExitBlock(ctx *BlockContext) {}

// EnterFunc_def is called when production func_def is entered.
func (s *BaseTurnkeyParserListener) EnterFunc_def(ctx *Func_defContext) {}

// ExitFunc_def is called when production func_def is exited.
func (s *BaseTurnkeyParserListener) ExitFunc_def(ctx *Func_defContext) {}

// EnterParameters is called when production parameters is entered.
func (s *BaseTurnkeyParserListener) EnterParameters(ctx *ParametersContext) {}

// ExitParameters is called when production parameters is exited.
func (s *BaseTurnkeyParserListener) ExitParameters(ctx *ParametersContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseTurnkeyParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseTurnkeyParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterInt_expression is called when production int_expression is entered.
func (s *BaseTurnkeyParserListener) EnterInt_expression(ctx *Int_expressionContext) {}

// ExitInt_expression is called when production int_expression is exited.
func (s *BaseTurnkeyParserListener) ExitInt_expression(ctx *Int_expressionContext) {}

// EnterFloat_expression is called when production float_expression is entered.
func (s *BaseTurnkeyParserListener) EnterFloat_expression(ctx *Float_expressionContext) {}

// ExitFloat_expression is called when production float_expression is exited.
func (s *BaseTurnkeyParserListener) ExitFloat_expression(ctx *Float_expressionContext) {}

// EnterBool_expression is called when production bool_expression is entered.
func (s *BaseTurnkeyParserListener) EnterBool_expression(ctx *Bool_expressionContext) {}

// ExitBool_expression is called when production bool_expression is exited.
func (s *BaseTurnkeyParserListener) ExitBool_expression(ctx *Bool_expressionContext) {}

// EnterString_expression is called when production string_expression is entered.
func (s *BaseTurnkeyParserListener) EnterString_expression(ctx *String_expressionContext) {}

// ExitString_expression is called when production string_expression is exited.
func (s *BaseTurnkeyParserListener) ExitString_expression(ctx *String_expressionContext) {}

// EnterIdent_expression is called when production ident_expression is entered.
func (s *BaseTurnkeyParserListener) EnterIdent_expression(ctx *Ident_expressionContext) {}

// ExitIdent_expression is called when production ident_expression is exited.
func (s *BaseTurnkeyParserListener) ExitIdent_expression(ctx *Ident_expressionContext) {}

// EnterCall_expression is called when production call_expression is entered.
func (s *BaseTurnkeyParserListener) EnterCall_expression(ctx *Call_expressionContext) {}

// ExitCall_expression is called when production call_expression is exited.
func (s *BaseTurnkeyParserListener) ExitCall_expression(ctx *Call_expressionContext) {}

// EnterTurn_expression is called when production turn_expression is entered.
func (s *BaseTurnkeyParserListener) EnterTurn_expression(ctx *Turn_expressionContext) {}

// ExitTurn_expression is called when production turn_expression is exited.
func (s *BaseTurnkeyParserListener) ExitTurn_expression(ctx *Turn_expressionContext) {}
