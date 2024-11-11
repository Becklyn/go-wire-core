package json

import (
	"runtime"

	"github.com/Becklyn/go-wire-core/env"
)

type Implementation interface {
	Marshal(v any) ([]byte, error)
	Unmarshal(data []byte, v any) error
	Valid(data []byte) bool
}

var jsonImplementation Implementation

func init() {
	if env.StringWithDefault("JSON_ENCODER", "std") == "std" {
		jsonImplementation = &StandardImplementation{}
		return
	}

	if runtime.GOARCH == "amd64" {
		jsonImplementation = &SonicImplementation{}
		return
	}

	jsonImplementation = &GoccyImplementation{}
}

func Marshal(v interface{}) ([]byte, error) {
	return jsonImplementation.Marshal(v)
}

func Unmarshal(data []byte, v interface{}) error {
	return jsonImplementation.Unmarshal(data, v)
}

func Valid(data []byte) bool {
	return jsonImplementation.Valid(data)
}
