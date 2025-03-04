package main 

type Err struct {
	line int
	token TokenKind
	msg string
}