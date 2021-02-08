// Code generated from Turnkey.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Turnkey

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTurnkeyListener is a complete listener for a parse tree produced by TurnkeyParser.
type BaseTurnkeyListener struct{}

var _ TurnkeyListener = &BaseTurnkeyListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTurnkeyListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTurnkeyListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTurnkeyListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTurnkeyListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseTurnkeyListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseTurnkeyListener) ExitStart(ctx *StartContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseTurnkeyListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseTurnkeyListener) ExitStatement(ctx *StatementContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseTurnkeyListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseTurnkeyListener) ExitBlock(ctx *BlockContext) {}

// EnterFunc_def is called when production func_def is entered.
func (s *BaseTurnkeyListener) EnterFunc_def(ctx *Func_defContext) {}

// ExitFunc_def is called when production func_def is exited.
func (s *BaseTurnkeyListener) ExitFunc_def(ctx *Func_defContext) {}

// EnterFunc_call is called when production func_call is entered.
func (s *BaseTurnkeyListener) EnterFunc_call(ctx *Func_callContext) {}

// ExitFunc_call is called when production func_call is exited.
func (s *BaseTurnkeyListener) ExitFunc_call(ctx *Func_callContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseTurnkeyListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseTurnkeyListener) ExitExpression(ctx *ExpressionContext) {}

// EnterInt_expression is called when production int_expression is entered.
func (s *BaseTurnkeyListener) EnterInt_expression(ctx *Int_expressionContext) {}

// ExitInt_expression is called when production int_expression is exited.
func (s *BaseTurnkeyListener) ExitInt_expression(ctx *Int_expressionContext) {}

// EnterFloat_expression is called when production float_expression is entered.
func (s *BaseTurnkeyListener) EnterFloat_expression(ctx *Float_expressionContext) {}

// ExitFloat_expression is called when production float_expression is exited.
func (s *BaseTurnkeyListener) ExitFloat_expression(ctx *Float_expressionContext) {}
