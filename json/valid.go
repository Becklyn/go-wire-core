package json

import (
	stdjson "encoding/json"
	"runtime"

	"github.com/Becklyn/go-wire-core/env"
	"github.com/bytedance/sonic"
	"github.com/goccy/go-json"
)

func Valid(data []byte) bool {
	if env.StringWithDefault("JSON_ENCODER", "std") == "std" {
		return stdjson.Valid(data)
	}

	if runtime.GOARCH == "amd64" {
		return sonic.Valid(data)
	}

	return json.Valid(data)
}
