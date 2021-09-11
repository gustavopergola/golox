package token

import "fmt"

type Token struct {
	LiteralValue interface{}
	Lexeme       string
	Type         TokenType
	Line         int
}

func (t *Token) ToString() string {
	return fmt.Sprintf("Token type: %s, lexeme: %s, line: %d", t.Type, t.Lexeme, t.Line)
}
