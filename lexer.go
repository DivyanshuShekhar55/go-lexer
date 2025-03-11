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
func (lex *Lexer) addToken(tokenType TokenKind, literal interface{}) {
	text := lex.source[lex.start:lex.curr]
	token := newToken(tokenType, text, literal, lex.line)
	lex.tokens = append(lex.tokens, token)
}

// returns curr char then moves ahead the curr pointer
func (lex *Lexer) advance() byte {
	ch :=  lex.charAt(lex.curr)
	lex.curr++
	return ch

}

// check what char is at a given position in the source file
func (lex *Lexer) charAt(pos int) byte {
	return lex.source[pos]
}

// we start the tokenisation process from here ...
func Tokenise(source string) []Token {
	lex := Lexer{}

	lex.start = 0
	lex.curr = 0
	lex.source = source
	lex.line = 1

	for !lex.isAtEnd() {
		//lex.advance()
		lex.start = lex.curr
		//lex.scanToken()

	}

	// we have reached the end of the file
	lex.addToken(EOF, nil)

	return lex.tokens
}

func (lex *Lexer) scanToken() {
	ch := lex.advance()
	// remember ch is the last char before moving

	switch ch {
	case '(':
		lex.addToken(LEFT_PAREN, nil)

	case ')':
		lex.addToken(RIGHT_PAREN, nil)

	case '{':
		lex.addToken(LEFT_BRACE, nil)

	case '}':
		lex.addToken(RIGHT_BRACE, nil)

	case ',':
		lex.addToken(COMMA, nil)

	case '.':
		lex.addToken(DOT, nil)

	case '-':
		lex.addToken(MINUS, nil)

	case '+':
		lex.addToken(PLUS, nil)

	case ';':
		lex.addToken(SEMICOLON, nil)

	case '*':
		lex.addToken(STAR, nil)

	case '!':
		var final TokenKind
		if lex.match('=') {
			final = BANG_EQUAL
		} else {
			final = BANG
		}
		lex.addToken(final, nil)

	case '=':
		var final TokenKind
		if lex.match('=') {
			final = EQUAL_EQUAL
		} else {
			final = EQUAL
		}
		lex.addToken(final, nil)

	case '<':
		var final TokenKind
		if lex.match('=') {
			final = LESS_EQUAL
		} else {
			final = LESS
		}
		lex.addToken(final, nil)

	case '>':
		var final TokenKind
		if lex.match('=') {
			final = GREATER_EQUAL
		} else {
			final = GREATER
		}
		lex.addToken(final, nil)

	case ' ':
	case '\r':
	case '\t':
		// Ignore whitespace
	case '\n':
		lex.line++

	case '/':
		// it may be a comment or div symbol, if comment ignore entire comment
		if lex.match('/') {
			for lex.peek() != '\n' && !lex.isAtEnd() {
				lex.advance()
			}
		} else {
			lex.addToken(SLASH, nil)
		}

	}
}

func (lex *Lexer) match(expected byte) bool {
	if lex.isAtEnd() {
		return false
	}
	if lex.charAt(lex.curr+1) != expected {
		return false
	}

	// matches case ... the true case
	lex.advance()
	return true
}

// looks at curr char
func (lex *Lexer) peek() byte {
	if lex.isAtEnd() {
		return '\x00'
	}
	return lex.charAt(lex.curr)
}

func (lex *Lexer) scanString(){
	// start is currently at '"', move the curr till we get another '"'
	for lex.peek() != '"' && !lex.isAtEnd(){
		if lex.peek() == '\n' {
			lex.line++
		}
		lex.advance()
	} 

	// say we reached EOF but no '"' was found, return error
	if lex.isAtEnd() {
		newErr(lex.line, STRING, "unterminated string found")
		return 
	}

	// no error, and curr is at the closing '"'

	// for value just terminate the starting and ending '"'
	value := lex.source[lex.start+1 : lex.curr-1]
	lex.addToken(STRING, value)
}
