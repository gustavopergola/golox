package scanner

import (
	"github.com/gustavopergola/golox/src/token"
)

type Scanner struct {
	SourceCode string
	tokens     []token.Token
	line       int
	start      int
	current    int
}

func (s *Scanner) ScanTokens() []token.Token {
	s.line = 1

	for _, r := range s.SourceCode {
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
	switch token.TokenType(r) {
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
	}
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
