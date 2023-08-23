package json

import (
	stdjson "encoding/json"
	"runtime"

	"github.com/Becklyn/go-wire-core/env"
	"github.com/bytedance/sonic"
	"github.com/goccy/go-json"
)

func Unmarshal(data []byte, v interface{}) error {
	if env.StringWithDefault("JSON_DECODER", "std") == "std" {
		return stdjson.Unmarshal(data, v)
	}

	if runtime.GOARCH == "amd64" {
		return sonic.Unmarshal(data, v)
	}

	return json.Unmarshal(data, v)
}
