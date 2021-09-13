package scanner

import (
	"fmt"

	"github.com/gustavopergola/golox/src/token"
)

type ScannerError struct {
	line    int
	message string
}

type Scanner struct {
	SourceCode string
	tokens     []token.Token
	Errors     []ScannerError
	line       int
	start      int
	current    int
}

func (s *Scanner) ScanTokens() []token.Token {
	s.line = 1

	for _, r := range s.SourceCode { // TODO: change this to manually control the reader
		if s.isAtEnd() {
			break
		}
		s.start = s.current
		s.scanToken(r)
	}

	s.addToken(token.EOF_TT)
	return s.tokens
}

func (s *Scanner) scanToken(r rune) {
	r_tt := token.TokenType(r)
	switch r_tt {
	case token.LEFT_PAREN_TT:
		s.addToken(token.LEFT_PAREN_TT)
	case token.RIGHT_PAREN_TT:
		s.addToken(token.RIGHT_PAREN_TT)
	case token.LEFT_BRACE_TT:
		s.addToken(token.LEFT_BRACE_TT)
	case token.RIGHT_BRACE_TT:
		s.addToken(token.RIGHT_BRACE_TT)
	case token.COMMA_TT:
		s.addToken(token.COMMA_TT)
	case token.DOT_TT:
		s.addToken(token.DOT_TT)
	case token.MINUS_TT:
		s.addToken(token.MINUS_TT)
	case token.PLUS_TT:
		s.addToken(token.PLUS_TT)
	case token.SEMICOLON_TT:
		s.addToken(token.SEMICOLON_TT)
	case token.STAR_TT:
		s.addToken(token.STAR_TT)
	default:
		s.addError(s.line, fmt.Sprintf("Unexpected charecter %s", string(r)))
	}
}

func (s *Scanner) addError(l int, msg string) {
	s.Errors = append(s.Errors, ScannerError{line: l, message: msg})
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.SourceCode)
}

func (s *Scanner) addToken(tt token.TokenType) {
	token := token.Token{
		Type:         tt,
		Lexeme:       string(tt),
		LiteralValue: nil,
		Line:         s.line,
	}

	s.tokens = append(s.tokens, token)
}
