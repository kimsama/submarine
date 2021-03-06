// This file was generated by typhen-api

package battle

import (
	"errors"
	"github.com/shiwano/submarine/server/battle/lib/typhenapi/core"
)

var _ = errors.New

// Pinger is a kind of TyphenAPI type.
type Pinger struct {
	ActorId    int64 `codec:"actor_id"`
	IsFinished bool  `codec:"is_finished"`
}

// Coerce the fields.
func (t *Pinger) Coerce() error {
	return nil
}

// Bytes creates the byte array.
func (t *Pinger) Bytes(serializer typhenapi.Serializer) ([]byte, error) {
	if err := t.Coerce(); err != nil {
		return nil, err
	}

	data, err := serializer.Serialize(t)
	if err != nil {
		return nil, err
	}

	return data, nil
}
