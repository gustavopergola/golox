package scanner

import (
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
	assert.Contains(t, errors[0].message, "invalid character")
}