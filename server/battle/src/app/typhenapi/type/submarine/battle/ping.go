// This file was generated by typhen-api

package battle

import (
	"errors"
)

var _ = errors.New

// Ping is kind of a TyphenAPI type.
type Ping struct {
	Message string `codec:"message"`
}

// Coerce the fields.
func (t *Ping) Coerce() error {
	return nil
}