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
	scanner := Scanner{
	}
	scanner.ScanTokens()
	errors := scanner.Errors
	assert.Len(t, errors, 0, "should not have errors")
	assert.Equal(t, scanner.tokens[0].Type, token.EOF_TT)
}

