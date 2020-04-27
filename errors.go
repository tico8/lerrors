package errors

import (
	"errors"
	"fmt"

	"golang.org/x/xerrors"
)

type err struct {
	label error
	err   error
	frame xerrors.Frame
}

// Error implementation of error interface
func (e *err) Error() string {
	return fmt.Sprintf("%+v", e.err)
}

// Unwrap implementation for errors.Unwrap
// https://golang.org/pkg/errors/#Unwrap
func (e *err) Unwrap() error {
	return errors.Unwrap(e.err)
}

// Is implementation for errors.Is
// https://golang.org/pkg/errors/#Is
func (e *err) Is(target error) bool {
	if e.label == target {
		return true
	}
	t, ok := target.(*err)
	if ok {
		return e.label == t.label
	}
	return false
}

// Format implementation of Formatter
// https://golang.org/pkg/fmt/#Formatter
func (e *err) Format(s fmt.State, v rune) {
	xerrors.FormatError(e, s, v)
}

// FormatError implementation of xerrors.Formatter
// https://godoc.org/golang.org/x/xerrors#Formatter
func (e *err) FormatError(p xerrors.Printer) (next error) {
	p.Print(e.err.Error())
	e.frame.Format(p)
	return nil
}

// Wrap create error with label
func Wrap(e, label error) error {
	return &err{
		label: label,
		err:   fmt.Errorf("%s: %w", label, e),
		frame: xerrors.Caller(1),
	}
}
