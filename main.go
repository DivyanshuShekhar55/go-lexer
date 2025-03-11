package main

import "fmt"

func main() {
	source := `while( a > b) {
	1 + 2
	}`
	tokens := Tokenise(source)
	for _, token := range tokens {
		fmt.Println(token)
	}
}
