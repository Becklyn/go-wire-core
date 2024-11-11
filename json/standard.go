package json

import "encoding/json"

type StandardImplementation struct{}

func (s StandardImplementation) Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func (s StandardImplementation) Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

func (s StandardImplementation) Valid(data []byte) bool {
	return json.Valid(data)
}
