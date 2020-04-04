package gomap

import (
	"fmt"
	"strings"
)

//WrongPathError raised when the path is not correct
type WrongPathError struct {
	path []string
}

func (e *WrongPathError) Error() string {
	return fmt.Sprintf("Wrong path %s", strings.Join(e.path, "/"))
}

//NewWrongPathError creates a new WrongPathError
func NewWrongPathError(path []string) *WrongPathError {
	return &WrongPathError{
		path: path,
	}
}

//WrongTypeError raised when the path is not correct
type WrongTypeError struct {
	expected string
	value    interface{}
}

func (e *WrongTypeError) Error() string {
	return fmt.Sprintf("Wrong type: expected %s actual %T", e.expected, e.value)
}

//NewWrongTypeError creates a new WrongPathError
func NewWrongTypeError(expected string, value interface{}) *WrongTypeError {
	return &WrongTypeError{
		expected: expected,
		value:    value,
	}
}
