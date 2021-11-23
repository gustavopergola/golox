package token

type TokenType string

const (
	// SINGLE RUNE TOKENS
	LEFT_PAREN_TT  TokenType = "("
	RIGHT_PAREN_TT TokenType = ")"
	LEFT_BRACE_TT  TokenType = "{"
	RIGHT_BRACE_TT TokenType = "}"
	COMMA_TT       TokenType = ","
	DOT_TT         TokenType = "."
	MINUS_TT       TokenType = "-"
	PLUS_TT        TokenType = "+"
	SEMICOLON_TT   TokenType = ";"
	STAR_TT        TokenType = "*"
	SLASH_TT       TokenType = "/"

	// ONE OR TWO RUNES TOKENS
	BANG_TT          TokenType = "!"
	BANG_EQUAL_TT    TokenType = "!="
	EQUAL_TT         TokenType = "="
	EQUAL_EQUAL_TT   TokenType = "=="
	GREATER_TT       TokenType = ">"
	GREATER_EQUAL_TT TokenType = ">="
	LESS_TT          TokenType = "<"
	LESS_EQUAL_TT    TokenType = "<="

	// LITERALS
	IDENTIFIER_TT TokenType = ""
	STRING_TT     TokenType = ""
	NUMBER_TT     TokenType = ""

	// KEYWORDS
	VAR_TT    TokenType = "var"
	FUN_TT    TokenType = "fun"
	AND_TT    TokenType = "and"
	CLASS_TT  TokenType = "class"
	ELSE_TT   TokenType = "else"
	FALSE_TT  TokenType = "false"
	FOR_TT    TokenType = "for"
	IF_TT     TokenType = "if"
	NIL_TT    TokenType = "nil"
	OR_TT     TokenType = "or"
	PRINT_TT  TokenType = "print"
	RETURN_TT TokenType = "return"
	SUPER_TT  TokenType = "super"
	THIS_TT   TokenType = "this"
	TRUE_TT   TokenType = "true"
	WHILE_TT  TokenType = "while"

	// FILE CONTROL
	EOF_TT     TokenType = "\000"
	NEWLINE_TT TokenType = "\n"
)

var ReservedWords = map[string]TokenType{
	"var":    VAR_TT,
	"fun":    FUN_TT,
	"and":    AND_TT,
	"class":  CLASS_TT,
	"else":   ELSE_TT,
	"false":  FALSE_TT,
	"for":    FOR_TT,
	"if":     IF_TT,
	"nil":    NIL_TT,
	"or":     OR_TT,
	"print":  PRINT_TT,
	"return": RETURN_TT,
	"super":  SUPER_TT,
	"this":   THIS_TT,
	"true":   TRUE_TT,
	"while":  WHILE_TT,
}
