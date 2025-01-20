package astprinter

import (
	"fmt"
	"strings"

	"github.com/brewinski/crafting-interpreters/pkg/expr"
)

type AstPrinter struct {
}

func (ap *AstPrinter) VisitBinaryExpr(binary expr.Binary[string]) string {
	return ap.parenthesize(binary.Operator.Lexeme, binary.Left, binary.Right)
}
func (ap *AstPrinter) VisitGroupingExpr(grouping expr.Grouping[string]) string {
	return ap.parenthesize("group", grouping.Expression)
}
func (ap *AstPrinter) VisitLiteralExpr(literal expr.Literal[string]) string {
	if literal.Value == nil {
		return "nil"
	}

	return string(fmt.Sprintf("%v", literal.Value))
}
func (ap *AstPrinter) VisitUnaryExpr(unary expr.Unary[string]) string {
	return ap.parenthesize(unary.Operator.Lexeme, unary.Right)
}

func (ap *AstPrinter) parenthesize(name string, exprs ...expr.Expr[string]) string {
	builder := strings.Builder{}

	builder.WriteString(fmt.Sprintf("(%s", name))
	for _, expr := range exprs {
		builder.WriteString(" ")
		builder.WriteString(expr.Accept(ap))
	}

	builder.WriteString(")")

	return builder.String()
}

func (ap *AstPrinter) Print(expr expr.Expr[string]) string {
	return expr.Accept(ap)
}
