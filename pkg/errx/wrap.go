package errx

import (
	"fmt"
	"github.com/pkg/errors"
)

type WrapError struct {
	code   string
	err    error
	detail interface{}
}

// Unwrap for errors.Is and errors.As
func (e *WrapError) Unwrap() error {
	return e.err
}

func (e *WrapError) Error() string {
	return e.err.Error()
}

func (e *WrapError) ErrorCode() string {
	return e.code
}

func (e *WrapError) WithDetail(detail interface{}) *WrapError {
	e.detail = detail
	return e
}

func (e *WrapError) Detail() interface{} {
	return e.detail
}

func New(code string, err interface{}) *WrapError {
	var e error
	if e1, ok := err.(error); ok {
		e = e1
	} else {
		e = errors.New(fmt.Sprint(err))
	}
	return &WrapError{
		code: code,
		err:  e,
	}
}

func As(err error) (*WrapError, bool) {
	target := &WrapError{}
	as := errors.As(err, &target)
	if as {
		return target, true
	}
	return nil, false
}
