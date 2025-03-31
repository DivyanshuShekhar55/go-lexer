package main

import (
	"fmt"
	lexer "github.com/DivyanshuShekhar55/go-lexer"
)

func main() {
	code := `while (x > 10) {
		print('10')
	}`
	fmt.Println(lexer.Tokenise(code))
}
