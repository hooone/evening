package errutil

import (
	"fmt"

	"golang.org/x/xerrors"
)

// Wrap is a simple wrapper around Errorf that is doing error wrapping.
func Wrap(message string, err error) error {
	if err == nil {
		return nil
	}

	return xerrors.Errorf("%v: %w", message, err)
}

// Wrapf is a simple wrapper around Errorf that is doing error wrapping
// Wrapf allows you to send a format and args instead of just a message.
func Wrapf(err error, message string, a ...interface{}) error {
	if err == nil {
		return nil
	}

	return Wrap(fmt.Sprintf(message, a...), err)
}
