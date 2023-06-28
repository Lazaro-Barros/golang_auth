package ops

import (
	"fmt"
	"runtime"
	"strconv"
)

// MyError é um tipo personalizado de erro que inclui informações de arquivo e linha.
type MyError struct {
	Inner          error
	Message        string
	HttpStatusCode int64
	Location       string
	Caller         int64
}

func (err *MyError) Error() string {
	return fmt.Sprintf("%v", err.Inner)
}

func NewErro(message string) error {
	_, file, line, _ := runtime.Caller(1)
	return &MyError{Inner: fmt.Errorf(message), Message: message, HttpStatusCode: 400, Location: file + ":" + strconv.Itoa(line)}
}

func Err(err error) error {
	_, ok := err.(*MyError)
	if ok {
		return err
	}
	_, file, line, _ := runtime.Caller(1)
	var myErr = MyError{
		Location:       file + ":" + strconv.Itoa(line),
		HttpStatusCode: 500,
		Inner:          err,
		Message:        err.Error(),
		Caller:         1,
	}

	return &myErr
}
