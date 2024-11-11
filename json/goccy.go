package json

import (
	goccy "github.com/goccy/go-json"
)

type GoccyImplementation struct{}

func (s GoccyImplementation) Marshal(v any) ([]byte, error) {
	return goccy.Marshal(v)
}

func (s GoccyImplementation) Unmarshal(data []byte, v any) error {
	return goccy.Unmarshal(data, v)
}

func (s GoccyImplementation) Valid(data []byte) bool {
	return goccy.Valid(data)
}
