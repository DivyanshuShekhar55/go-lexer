package main

// we need col and row of current token to scan it
// row = line and col = (start, curr)
// if token = While then start = curr = W and keep moving curr till we finish to token
// once curr = 'e' add token 'While' and skip start to character after 'e'

type Lexer struct {
	source string
	tokens []Token
	start  int
	curr   int
	line   int
}

// returns whether we have reached the EOF or not
func (lex *Lexer) isAtEnd() bool {
	return lex.curr >= len(lex.source)
}

// add a scanned token
func (lex *Lexer) addToken(token Token) {
	lex.tokens = append(lex.tokens, token)
}

// move the curr pointer
func (lex *Lexer) advance() {
	lex.curr++
}

// check what char is at a given position in the source file
func (lex *Lexer) charAt(pos int) byte {
	return lex.source[lex.curr]
}
