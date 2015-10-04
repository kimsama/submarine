// This file was generated by typhen-api

package submarine

import (
	"errors"
)

// PingMessage is kind of a TyphenAPI type.
type PingMessage struct {
	Message string `codec:"message"`
}

// MessageType returns message type.
func (e *PingMessage) MessageType() int32 {
	return 1
}

// Coerce fields.
func (e *PingMessage) Coerce() error {
	if e.Message == "" {
		return errors.New("Message is empty")
	}

	return nil
}
