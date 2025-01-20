package main

import (
	"fmt"

	astprinter "github.com/brewinski/crafting-interpreters/pkg/ast-printer"
	"github.com/brewinski/crafting-interpreters/pkg/expr"
	"github.com/brewinski/crafting-interpreters/pkg/token"
)

func main() {
	expr := expr.NewBinary(
		expr.NewUnary(
			token.NewToken(token.MINUS, "-", 1, nil),
			expr.NewLiteral[string](123),
		),
		token.NewToken(token.STAR, "*", 1, nil),
		expr.NewGrouping(
			expr.NewLiteral[string](47.5),
		),
	)

	astPrinter := astprinter.AstPrinter{}

	fmt.Println(astPrinter.Print(expr))
}
