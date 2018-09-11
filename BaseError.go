package commons

import (
	"fmt"
	"runtime"
)

// BaseErrorCat category of errors represented by BaseError.
type BaseErrorCat string

const (
	// ErrDecline the underlying operation has been declined.
	ErrDecline BaseErrorCat = "ErrDecline"

	// ErrFailure the underlying operation has generade a recoverable failure.
	ErrFailure BaseErrorCat = "ErrFailure"

	// ErrFatal the underlying operation has generade a fatal failure.
	ErrFatal BaseErrorCat = "ErrFatal"
)

// BaseError base of all errors.
type BaseError struct {
	Cat  BaseErrorCat
	Rsn  string
	Msg  string
	Func string
	Line int
	File string
}

func (e BaseError) Error() string {
	return fmt.Sprintf(
		"%s: %s -- %s (%s\\@%s:%d)",
		e.Cat, e.Rsn, e.Msg, e.Func, e.File, e.Line)
}

// NewError instantiate a new error.
func NewError(cat BaseErrorCat, rsn, msg string) (err *BaseError) {
	pc, file, line, _ := runtime.Caller(1)

	err = &BaseError{
		Cat:  cat,
		Rsn:  rsn,
		Msg:  msg,
		Line: line,
		File: file,
		Func: runtime.FuncForPC(pc).Name()}

	return
}

// NewDecline instantiate a new error.
func NewDecline(rsn, msg string) (err *BaseError) {
	return NewError(ErrDecline, rsn, msg)
}

// NewFailure instantiate a new error.
func NewFailure(rsn, msg string) (err *BaseError) {
	return NewError(ErrFailure, rsn, msg)
}

// NewFatal instantiate a new error.
func NewFatal(rsn, msg string) (err *BaseError) {
	return NewError(ErrFatal, rsn, msg)
}
