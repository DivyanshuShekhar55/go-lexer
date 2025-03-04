package main

// we need col and row of current token to scan it
// row = line and col = (start, curr)
// if token = While then start = curr = W and keep moving curr till we finish to token
// once curr = 'e' add token 'While' and skip start to character after 'e'

type Lexer struct{
	source string
	tokens []Token
	start int 
	curr int
	line int
}
