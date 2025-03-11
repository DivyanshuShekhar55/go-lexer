package main

import "fmt"

type TokenKind int

const (

	// iota means declaration of enum (starts with 0)
	// single character tokens
	LEFT_PAREN TokenKind = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR

	// one or two character tokens
	// bang means '!', bang_equal is '!='
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	// literals
	IDENTIFIER
	STRING
	NUMBER

	// keywords
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE

	// end of file
	EOF

	// for general errors where unexpected tokens are provided
	UNKNOWN
)

type Token struct {
	tokenType TokenKind
	lexeme    string
	literal   interface{} // can hold any type
	line      int
}

func newToken(tokenType TokenKind, lexeme string, literal interface{}, line int) Token {

	return Token{
		tokenType: tokenType,
		lexeme:    lexeme,
		literal:   literal,
		line:      line,
	}
}

func (t *Token) toString() string {
	return fmt.Sprintf("%s %s %s", t.tokenType, t.lexeme, t.literal)
}

var keywords = map[string]TokenKind{
	"and":    AND,
	"class":  CLASS,
	"else":   ELSE,
	"false":  FALSE,
	"for":    FOR,
	"fun":    FUN,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"print":  PRINT,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
	"true":   TRUE,
	"var":    VAR,
	"while":  WHILE,
}
