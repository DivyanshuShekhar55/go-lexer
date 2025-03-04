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

// move the curr pointer
func (lex *Lexer) advance() {
	lex.curr++
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
	lex.advance()
	ch := lex.charAt(lex.curr)

	switch ch {
	case '(':
		lex.addToken(LEFT_PAREN, nil)
		break
	case ')':
		lex.addToken(RIGHT_PAREN, nil)
		break
	case '{':
		lex.addToken(LEFT_BRACE, nil)
		break
	case '}':
		lex.addToken(RIGHT_BRACE, nil)
		break
	case ',':
		lex.addToken(COMMA, nil)
		break
	case '.':
		lex.addToken(DOT, nil)
		break
	case '-':
		lex.addToken(MINUS, nil)
		break
	case '+':
		lex.addToken(PLUS, nil)
		break
	case ';':
		lex.addToken(SEMICOLON, nil)
		break
	case '*':
		lex.addToken(STAR, nil)
		break
	case '!':
		var final TokenKind
		if lex.match('=') {
			final = BANG_EQUAL
		} else {
			final = BANG
		}
		lex.addToken(final, nil)
		break
	case '=':
		var final TokenKind
		if lex.match('=') {
			final = EQUAL_EQUAL
		} else {
			final = EQUAL
		}
		lex.addToken(final, nil)
		break
	case '<':
		var final TokenKind
		if lex.match('=') {
			final = LESS_EQUAL
		} else {
			final = LESS
		}
		lex.addToken(final, nil)
		break
	case '>':
		var final TokenKind
		if lex.match('=') {
			final = GREATER_EQUAL
		} else {
			final = GREATER
		}
		lex.addToken(final, nil)
		break
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
