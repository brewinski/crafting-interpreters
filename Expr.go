package main

import "github.com/brewinski/crafting-interpreters/pkg/token"

type Expr 

type Visitor[R any] interface {
	VisitBinaryExpr(binary Binary[R]) R
	VisitGroupingExpr(grouping Grouping[R]) R
	VisitLiteralExpr(literal Literal[R]) R
	VisitUnaryExpr(unary Unary[R]) R
}

type Binary[R any] struct {
	left     Expr
	operator token.Token
	right    Expr
}

func NewBinary[R any](left Expr, operator token.Token, right Expr) Binary[R] {
	return Binary[R]{
		left,
		operator,
		right,
	}
}

func (binary Binary[R]) Accept(visitor Visitor[R]) R {
	return visitor.VisitBinaryExpr(binary)
}

type Grouping[R any] struct {
	expression Expr
}

func NewGrouping[R any](expression Expr) Grouping[R] {
	return Grouping[R]{
		expression,
	}
}

func (grouping Grouping[R]) Accept(visitor Visitor[R]) R {
	return visitor.VisitGroupingExpr(grouping)
}

type Literal[R any] struct {
	value interface{}
}

func NewLiteral[R any](value interface{}) Literal[R] {
	return Literal[R]{
		value,
	}
}

func (literal Literal[R]) Accept(visitor Visitor[R]) R {
	return visitor.VisitLiteralExpr(literal)
}

type Unary[R any] struct {
	operator token.Token
	right    Expr
}

func NewUnary[R any](operator token.Token, right Expr) Unary[R] {
	return Unary[R]{
		operator,
		right,
	}
}

func (unary Unary[R]) Accept(visitor Visitor[R]) R {
	return visitor.VisitUnaryExpr(unary)
}
