// Code generated from Turnkey.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Turnkey

import "github.com/antlr/antlr4/runtime/Go/antlr"

// TurnkeyListener is a complete listener for a parse tree produced by TurnkeyParser.
type TurnkeyListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterInt_expression is called when entering the int_expression production.
	EnterInt_expression(c *Int_expressionContext)

	// EnterFloat_expression is called when entering the float_expression production.
	EnterFloat_expression(c *Float_expressionContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitInt_expression is called when exiting the int_expression production.
	ExitInt_expression(c *Int_expressionContext)

	// ExitFloat_expression is called when exiting the float_expression production.
	ExitFloat_expression(c *Float_expressionContext)
}
