# Logging

We provide a preconfigured logging service (golog).

```go
logger := logging.New(env.New())
```

The log level can be configured using the `LOG_LEVEL` env variable. Default level is `info`.

Possible values:

- `debug`
- `info` (this is the default / fallback level)
- `warn`
- `error`
- `fatal`

The log format can be configured using the `LOG_FORMAT` env variable. Default is `console`.

Possible values:

- `console` (this is the default / fallback format)
- `json`
