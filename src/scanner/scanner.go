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
	Errors     []ScannerError

	runes  []rune
	tokens []token.Token

	line    int
	start   int
	current int
}

func (s *Scanner) ScanTokens() []token.Token {
	s.line = 1
	s.runes = []rune(s.SourceCode)

	for {
		if s.isAtEnd() {
			break
		}
		s.start = s.current
		s.scanToken()
	}

	s.addToken(token.EOF_TT)
	return s.tokens
}

func (s *Scanner) scanToken() {
	var (
		ttToAdd  token.TokenType
		hadError bool
	)
	r := s.currentRune()
	s.advance()
	r_tt := token.TokenType(r)

	switch r_tt {
	// one rune
	case token.LEFT_PAREN_TT:
		ttToAdd = token.LEFT_PAREN_TT
	case token.RIGHT_PAREN_TT:
		ttToAdd = token.RIGHT_PAREN_TT
	case token.LEFT_BRACE_TT:
		ttToAdd = token.LEFT_BRACE_TT
	case token.RIGHT_BRACE_TT:
		ttToAdd = token.RIGHT_BRACE_TT
	case token.COMMA_TT:
		ttToAdd = token.COMMA_TT
	case token.DOT_TT:
		ttToAdd = token.DOT_TT
	case token.MINUS_TT:
		ttToAdd = token.MINUS_TT
	case token.PLUS_TT:
		ttToAdd = token.PLUS_TT
	case token.SEMICOLON_TT:
		ttToAdd = token.SEMICOLON_TT
	case token.STAR_TT:
		ttToAdd = token.STAR_TT

	// one or two runes
	case token.BANG_TT:
		ttToAdd = token.BANG_TT
		if s.match(token.EQUAL_TT) {
			ttToAdd = token.BANG_EQUAL_TT
		}
	case token.EQUAL_TT:
		ttToAdd = token.EQUAL_TT
		if s.match(token.EQUAL_TT) {
			ttToAdd = token.EQUAL_EQUAL_TT
		}
	case token.LESS_TT:
		ttToAdd = token.LESS_TT
		if s.match(token.EQUAL_TT) {
			ttToAdd = token.LESS_EQUAL_TT
		}
	case token.GREATER_TT:
		ttToAdd = token.GREATER_TT
		if s.match(token.EQUAL_TT) {
			ttToAdd = token.GREATER_EQUAL_TT
		}
	default:
		s.addError(s.line, fmt.Sprintf("Unexpected character %s", string(r)))
		hadError = true
	}

	if !hadError {
		s.addToken(ttToAdd)
	}
}

func (s *Scanner) match(expected token.TokenType) bool {
	if s.isAtEnd() || token.TokenType(s.currentRune()) != expected {
		return false
	}

	s.current += 1
	return true
}

func (s *Scanner) addError(l int, msg string) {
	s.Errors = append(s.Errors, ScannerError{line: l, message: msg})
}

func (s *Scanner) addToken(tt token.TokenType) {
	t := token.Token{
		Type:         tt,
		Lexeme:       string(tt),
		LiteralValue: nil,
		Line:         s.line,
	}

	s.tokens = append(s.tokens, t)
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.SourceCode)
}

func (s *Scanner) currentRune() rune {
	return s.runes[s.current]
}

func (s *Scanner) advance() {
	s.current += 1
}
