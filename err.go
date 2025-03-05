package main

import "fmt"

type Err struct {
	line  int
	token TokenKind
	msg   string
}

// now our Err struct implements the error interface of Go, by using the Error method
// so any value based on Err struct will be auto reognised as an error by Go
func (e *Err) Error() string {
	return fmt.Sprintf("Err at %d: %d \n%s", e.line, e.token, e.msg)
}

// var FLAG_STATUS int8 = 0, don't feel quite necessary also unsafe and arbitrary updates can be there

// this is the Error Generation Part :
func newErr(line int, token TokenKind, msg string) error {

	err := &Err{
		line:  line,
		token: token,
		msg:   msg,
	}

	//FLAG_STATUS = 1
	// err automatically recognised as error by Golang
	return err

}

// error reporting part :
func (err *Err) reportErr() {
	// the Error() func will be called automatically here as err has error type
	// anywhere this err value is used as string, it calls the Error() method
	// it could be sending to UI (via API response), file logging, console logging (logging a string)
	// in API response it may be like return APIResponse{Success: false, Message: err.Error()},
	// as API returns strings only so err.Error() needs be called which returns string error
	fmt.Println(err)
}
