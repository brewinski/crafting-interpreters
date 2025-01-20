package scanner

import (
	"fmt"
	"unicode"

	interr "github.com/brewinski/crafting-interpreters/pkg/error"
	"github.com/brewinski/crafting-interpreters/pkg/token"
)

type Scanner struct {
	source  string
	tokens  []token.Token
	start   int
	current int
	line    int
	col     int
}

func NewScanner(source string) Scanner {
	return Scanner{
		source:  source,
		tokens:  []token.Token{},
		start:   0,
		current: 0,
		line:    1,
		col:     0,
	}
}

func (s *Scanner) ScanTokens() []token.Token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}

	s.tokens = append(s.tokens, token.NewToken(token.EOF, "", s.line, ""))

	return s.tokens
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) scanToken() {
	c := s.advance()
	switch c {
	case '(':
		s.addToken(token.LEFT_PAREN, "")
	case ')':
		s.addToken(token.RIGHT_PAREN, "")
	case '{':
		s.addToken(token.LEFT_BRACE, "")
	case '}':
		s.addToken(token.RIGHT_BRACE, "")
	case ',':
		s.addToken(token.COMMA, "")
	case '.':
		s.addToken(token.DOT, "")
	case '-':
		s.addToken(token.MINUS, "")
	case '+':
		s.addToken(token.PLUS, "")
	case ';':
		s.addToken(token.SEMICOLON, "")
	case '*':
		s.addToken(token.STAR, "")
	case '!':
		if s.match('=') {
			s.addToken(token.BANG_EQUAL, "")
			break
		}
		s.addToken(token.BANG, "")
	case '=':
		if s.match('=') {
			s.addToken(token.EQUAL_EQUAL, "")
			break
		}
		s.addToken(token.EQUAL, "")
	case '>':
		if s.match('=') {
			s.addToken(token.GREATER_EQUAL, "")
			break
		}
		s.addToken(token.GREATER, "")
	case '<':
		if s.match('=') {
			s.addToken(token.LESS_EQUAL, "")
			break
		}

		s.addToken(token.LESS, "")
	case '/':
		if s.match('/') {
			// a comment does until the end of the line.
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		}
		s.addToken(token.SLASH, "")
	case ' ', '\t', '\r':
		break
	case '\n':
		s.line++
		s.col = 0
	case '"':
		s.string()
	default:
		if s.isNumber(c) {
			s.number()
			break
		}
		if s.isAlpha(c) {
			s.identifier()
			break
		}
		interr.Error(s.line, s.col, fmt.Sprintf("Unexpected character: %s", string(c)))
	}
}

func (s *Scanner) addToken(tokenType token.TokenType, literal any) {
	text := s.source[s.start:s.current]

	s.tokens = append(s.tokens, token.NewToken(tokenType, text, s.line, literal))
}

func (s *Scanner) advance() byte {
	c := s.source[s.current]
	s.current++
	s.col++

	return c
}

func (s *Scanner) match(expected byte) bool {
	if s.isAtEnd() {
		return false
	}

	if s.source[s.current] != expected {
		return false
	}

	s.current++
	s.col++
	return true
}

func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		return 0
	}

	return s.source[s.current]
}

func (s *Scanner) string() {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
			s.col = 0
		}
		s.advance()
	}

	if s.isAtEnd() {
		interr.Error(s.line, s.col, "Unterminated string.")
		return
	}

	s.advance()

	// get the string excluding the leading and trailing "
	stringValue := s.source[s.start+1 : s.current-1]

	s.addToken(token.STRING, stringValue)
}

func (s *Scanner) number() {
	for s.isNumber(s.peek()) {
		s.advance()
	}

	// look for fractional number
	if s.peek() == '.' && s.isNumber(s.peekNext()) {
		s.advance()
		for s.isNumber(s.peek()) {
			s.advance()
		}
	}

	// get the string excluding the leading and trailing "
	stringValue := s.source[s.start:s.current]

	s.addToken(token.NUMBER, stringValue)
}

func (s *Scanner) peekNext() byte {
	if s.current+1 >= len(s.source) {
		return 0
	}

	return s.source[s.current+1]
}

func (s *Scanner) identifier() {
	for s.isAlpha(s.peek()) {
		s.advance()
	}

	text := s.source[s.start:s.current]

	tokenType, ok := token.Keywords[text]
	if !ok {
		tokenType = token.IDENTIFIER
	}

	s.addToken(tokenType, "")
}

func (s *Scanner) isAlpha(r byte) bool {
	return unicode.IsLetter(rune(r)) || r == '_'
}

func (s *Scanner) isNumber(r byte) bool {
	return unicode.IsDigit(rune(r))
}

func (s *Scanner) isAlphaNumeric(r byte) bool {
	return s.isAlpha(r) || s.isNumber(r)
}
