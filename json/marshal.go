package json

import (
	stdjson "encoding/json"
	"runtime"

	"github.com/Becklyn/go-wire-core/env"
	"github.com/bytedance/sonic"
	"github.com/goccy/go-json"
)

func Marshal(v interface{}) ([]byte, error) {
	if env.StringWithDefault("JSON_ENCODER", "std") == "std" {
		return stdjson.Marshal(v)
	}

	if runtime.GOARCH == "amd64" {
		return sonic.Marshal(v)
	}

	return json.Marshal(v)
}
