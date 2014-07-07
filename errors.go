package trailer

import (
	"fmt"
)

type TrailerError struct {
	str string
}

func Errorf(format string, a ...interface{}) error {
	return &TrailerError{
		str: fmt.Sprintf(format, a...),
	}
}

func (err *TrailerError) Error() string {
	return err.str
}
