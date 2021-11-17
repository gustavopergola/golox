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

	s.addToken(token.EOF_TT, nil)
	return s.tokens
}

func (s *Scanner) scanToken() {
	var hadError bool

	ttToAdd := token.EOF_TT

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
	case token.SLASH_TT:
		// Check for comments in line
		if s.match(token.SLASH_TT) {
			for s.peek() != token.NEWLINE_TT && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(token.SLASH_TT, nil)
		}
	case " ":
		break
	case "\r":
		break
	case "\t":
		break
	case "\"":
		s.scanString()
	case token.NEWLINE_TT:
		s.line++
		break
	default:
		s.addError(s.line, fmt.Sprintf("Unexpected character %s", string(r)))
		hadError = true
	}

	if !hadError && ttToAdd != token.EOF_TT {
		s.addToken(ttToAdd, nil)
	}
}

// Checks if current rune matches expected param. Advance if it does.
// Could be read as "conditional advance"
func (s *Scanner) match(expected token.TokenType) bool {
	if lookahead := s.peek(); lookahead != expected {
		return false
	}

	s.current += 1
	return true
}

// Look into current without advancing, return as token type for ease of use.
func (s *Scanner) peek() token.TokenType {
	if s.isAtEnd() {
		return token.EOF_TT
	}
	return token.TokenType(s.currentRune())
}

// Simply append an Error to the scanner.
func (s *Scanner) addError(l int, msg string) {
	s.Errors = append(s.Errors, ScannerError{line: l, message: msg})
}

func (s *Scanner) addToken(tt token.TokenType, value interface{}) {
	t := token.Token{
		Type:         tt,
		Lexeme:       string(tt),
		LiteralValue: value,
		Line:         s.line,
	}

	s.tokens = append(s.tokens, t)
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.SourceCode)
}

// Auxiliary to get current rune.
func (s *Scanner) currentRune() rune {
	return s.runes[s.current]
}

func (s *Scanner) advance() {
	s.current += 1
}


func (s *Scanner) scanString() {
	current := s.peek()
	for current != "\"" && s.isAtEnd() {
		if current == token.NEWLINE_TT {
			s.line++
		}
		s.advance()
	}

	if s.isAtEnd() {
		s.addError(s.line, "unterminated string.")
	}

	//terminating \"
	s.advance()

	value := s.SourceCode[s.start + 1:s.current - 1]
	s.addToken(token.STRING_TT, value)
}

func (s *Scanner) PrintTokens() {
	for _, t := range s.tokens {
		fmt.Println(t.ToString())
	}
}