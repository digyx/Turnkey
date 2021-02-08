// Code generated from Turnkey.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Turnkey

import "github.com/antlr/antlr4/runtime/Go/antlr"

// TurnkeyListener is a complete listener for a parse tree produced by TurnkeyParser.
type TurnkeyListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterFunc_def is called when entering the func_def production.
	EnterFunc_def(c *Func_defContext)

	// EnterFunc_call is called when entering the func_call production.
	EnterFunc_call(c *Func_callContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterInt_expression is called when entering the int_expression production.
	EnterInt_expression(c *Int_expressionContext)

	// EnterFloat_expression is called when entering the float_expression production.
	EnterFloat_expression(c *Float_expressionContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitFunc_def is called when exiting the func_def production.
	ExitFunc_def(c *Func_defContext)

	// ExitFunc_call is called when exiting the func_call production.
	ExitFunc_call(c *Func_callContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitInt_expression is called when exiting the int_expression production.
	ExitInt_expression(c *Int_expressionContext)

	// ExitFloat_expression is called when exiting the float_expression production.
	ExitFloat_expression(c *Float_expressionContext)
}
