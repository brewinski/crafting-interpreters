package token

import "fmt"

type TokenType int

const (
	// single-character tokens.
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR

	// One or two character tokens.
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	// Literals
	IDENTIFIER
	STRING
	NUMBER

	// Keywords
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE

	EOF
)

var (
	Keywords = map[string]TokenType{
		"and":    AND,
		"class":  CLASS,
		"else":   ELSE,
		"false":  FALSE,
		"for":    FOR,
		"fun":    FUN,
		"if":     IF,
		"nil":    NIL,
		"or":     OR,
		"print":  PRINT,
		"return": RETURN,
		"super":  SUPER,
		"this":   THIS,
		"true":   TRUE,
		"var":    VAR,
		"while":  WHILE,
	}
)

func (tt TokenType) String() string {
	switch tt {
	case LEFT_PAREN:
		return "("
	case RIGHT_PAREN:
		return ")"
	case LEFT_BRACE:
		return "{"
	case RIGHT_BRACE:
		return "}"
	case COMMA:
		return ","
	case DOT:
		return "."
	case MINUS:
		return "-"
	case PLUS:
		return "+"
	case SEMICOLON:
		return ";"
	case SLASH:
		return "/"
	case STAR:
		return "*"
	case EOF:
		return "EOF"
	case STRING:
		return "STRING"
	case BANG:
		return "!"
	case BANG_EQUAL:
		return "!="
	case NUMBER:
		return "NUMBER"
	case IDENTIFIER:
		return "IDENTIFIER"
	case EQUAL:
		return "="
	case EQUAL_EQUAL:
		return "=="
	case GREATER:
		return ">"
	case GREATER_EQUAL:
		return ">="
	case LESS:
		return "<"
	case LESS_EQUAL:
		return "<="
	default:
		for k, v := range Keywords {
			if tt == v {
				return k
			}
		}
		return "UNKNOWN"
	}
}

type Token struct {
	TokenType TokenType
	Lexeme    string
	Line      int
	Literal   any
}

func NewToken(tokenType TokenType, lexeme string, line int, literal any) Token {
	return Token{tokenType, lexeme, line, literal}
}

func (t *Token) String() string {
	return fmt.Sprintf("%s %s %s", t.TokenType, t.Lexeme, t.Literal)
}
