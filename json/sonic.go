package json

import "github.com/bytedance/sonic"

type SonicImplementation struct{}

func (s SonicImplementation) Marshal(v any) ([]byte, error) {
	return sonic.Marshal(v)
}

func (s SonicImplementation) Unmarshal(data []byte, v any) error {
	return sonic.Unmarshal(data, v)
}

func (s SonicImplementation) Valid(data []byte) bool {
	return sonic.Valid(data)
}
