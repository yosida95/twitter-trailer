package trailer

import (
	"fmt"
)

type TrailerError struct {
	str string
}

func errorf(format string, a ...interface{}) *TrailerError {
	return &TrailerError{
		str: fmt.Sprintf(format, a...),
	}
}

func (err *TrailerError) Error() string {
	return err.str
}
