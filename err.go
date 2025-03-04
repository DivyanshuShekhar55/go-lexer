package main

import "fmt"

type Err struct {
	line  int
	token TokenKind
	msg   string
}

func (e *Err) reportErr(line int, token TokenKind, msg string) error {
	return fmt.Errorf("Err at %d : %s \n%s", line, token, msg) 
}

