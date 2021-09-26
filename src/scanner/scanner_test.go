package scanner

import (
	"github.com/gustavopergola/golox/src/token"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScanner_InvalidCharacter(t *testing.T) {
	invalidCharacter := "@"
	scanner := Scanner{
		SourceCode: invalidCharacter,
	}
	scanner.ScanTokens()
	errors := scanner.Errors
	assert.Len(t, errors, 1, "should have 1 error")
	assert.Contains(t, errors[0].message, invalidCharacter)
	assert.Contains(t, errors[0].message, "Unexpected character")
}

func TestScanner_AddEOFToken(t *testing.T) {
	scanner := Scanner{}
	scanner.ScanTokens()
	errors := scanner.Errors
	assert.Len(t, errors, 0, "should not have errors")
	assert.Equal(t, scanner.tokens[0].Type, token.EOF_TT, "EOF should be added")
}

func TestScanner_AddOneRuneToken(t *testing.T) {
	semiColon := string(token.SEMICOLON_TT)
	scanner := Scanner{
		SourceCode: semiColon,
	}
	scanner.ScanTokens()
	errors := scanner.Errors
	assert.Len(t, errors, 0, "should not have errors")
	assert.Equal(t, token.SEMICOLON_TT, scanner.tokens[0].Type, "semicolon should have been added")
	assert.Len(t, scanner.tokens, 2, "should have semicolon and EOF")
}

func TestScanner_AddOneOrTwoRuneToken(t *testing.T) {
	greaterOrEqual := string(token.GREATER_EQUAL_TT)
	scanner := Scanner{
		SourceCode: greaterOrEqual,
	}
	scanner.ScanTokens()
	errors := scanner.Errors
	assert.Len(t, errors, 0, "should not have errors")
	assert.Equal(t, token.GREATER_EQUAL_TT, scanner.tokens[0].Type, "greater_or_equal should have been added")
	assert.Len(t, scanner.tokens, 2, "should have greater_or_equal and EOF")
}
