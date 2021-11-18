package scanner

import (
	"github.com/gustavopergola/golox/src/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

func scannerAsserts(t *testing.T, code string, nTokens, lines int) *Scanner {
	s := Scanner{SourceCode: code}
	s.ScanTokens()
	assert.Len(t, s.Errors, 0, "should not have errors")
	assert.Len(t, s.tokens, nTokens)
	assert.Equal(t, lines, s.line)
	return &s
}

func TestScanner_AddEOFToken(t *testing.T) {
	s := scannerAsserts(t, "", 1, 1)
	assert.Equal(t, s.tokens[0].Type, token.EOF_TT, "EOF should be added")
}

func TestScanner_AddOneRuneToken(t *testing.T) {
	s := scannerAsserts(t, string(token.SEMICOLON_TT), 2, 1)
	assert.Equal(t, token.SEMICOLON_TT, s.tokens[0].Type, "semicolon should have been added")
}

func TestScanner_AddOneOrTwoRuneToken(t *testing.T) {
	s := scannerAsserts(t, string(token.GREATER_EQUAL_TT), 2, 1)
	assert.Equal(t, token.GREATER_EQUAL_TT, s.tokens[0].Type, "greater_or_equal should have been added")
}

func TestScanner_IgnoresComments(t *testing.T) {
	scannerAsserts(t, "// this is a coment!", 1, 1)
}

func TestScanner_IgnoresMeaningless(t *testing.T) {
	scannerAsserts(t, "\t\r", 1, 1)
}

func TestScanner_AddNewLine(t *testing.T) {
	scannerAsserts(t, "\n", 1, 2)
}

func TestScanner_ScanString(t *testing.T) {
	s := scannerAsserts(t, "\"foobar\n\"", 2, 2)
	assert.Equal(t, s.tokens[0].LiteralValue, "foobar\n")
}

func TestScanner_ScanEmptyString(t *testing.T) {
	s := scannerAsserts(t, "\"\"", 2, 1)
	assert.Equal(t, s.tokens[0].LiteralValue, "")
}

func TestScanner_InvalidCharacter(t *testing.T) {
	invalidCharacter := "@"
	scanner := Scanner{
		SourceCode: invalidCharacter,
	}
	scanner.ScanTokens()
	errors := scanner.Errors
	assert.Len(t, errors, 1, "should have 1 error")
	assert.Contains(t, errors[0].message, invalidCharacter)
	assert.Contains(t, errors[0].message, "unexpected character")
}

func TestScanner_InvalidCharacterAfterMatch(t *testing.T) {
	invalidCharacter := "@"
	scanner := Scanner{
		SourceCode: "<" + invalidCharacter,
	}
	scanner.ScanTokens()
	errors := scanner.Errors
	assert.Len(t, errors, 1, "should have 1 error")
	assert.Contains(t, errors[0].message, invalidCharacter)
	assert.Contains(t, errors[0].message, "unexpected character")
}

func TestScanner_UnterminatedString(t *testing.T) {
	unterminatedString := "\"foobarz"
	scanner := Scanner{
		SourceCode: unterminatedString,
	}
	scanner.ScanTokens()
	errors := scanner.Errors
	assert.Len(t, errors, 1, "should have 1 error")
	assert.Contains(t, errors[0].message, "unterminated string")
}