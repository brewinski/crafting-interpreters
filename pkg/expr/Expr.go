package expr

import "github.com/brewinski/crafting-interpreters/pkg/token"

type Expr[R any] interface {
	Accept(visitor Visitor[R]) R
}

type Visitor[R any] interface {
	VisitBinaryExpr(binary Binary[R]) R
	VisitGroupingExpr(grouping Grouping[R]) R
	VisitLiteralExpr(literal Literal[R]) R
	VisitUnaryExpr(unary Unary[R]) R
}

type Binary[R any] struct {
	Left     Expr[R]
	Operator token.Token
	Right    Expr[R]
}

func NewBinary[R any](Left Expr[R], Operator token.Token, Right Expr[R]) Binary[R] {
	return Binary[R]{
		Left,
		Operator,
		Right,
	}
}

func (binary Binary[R]) Accept(visitor Visitor[R]) R {
	return visitor.VisitBinaryExpr(binary)
}

type Grouping[R any] struct {
	Expression Expr[R]
}

func NewGrouping[R any](Expression Expr[R]) Grouping[R] {
	return Grouping[R]{
		Expression,
	}
}

func (grouping Grouping[R]) Accept(visitor Visitor[R]) R {
	return visitor.VisitGroupingExpr(grouping)
}

type Literal[R any] struct {
	Value interface{}
}

func NewLiteral[R any](Value interface{}) Literal[R] {
	return Literal[R]{
		Value,
	}
}

func (literal Literal[R]) Accept(visitor Visitor[R]) R {
	return visitor.VisitLiteralExpr(literal)
}

type Unary[R any] struct {
	Operator token.Token
	Right    Expr[R]
}

func NewUnary[R any](Operator token.Token, Right Expr[R]) Unary[R] {
	return Unary[R]{
		Operator,
		Right,
	}
}

func (unary Unary[R]) Accept(visitor Visitor[R]) R {
	return visitor.VisitUnaryExpr(unary)
}
