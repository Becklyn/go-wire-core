# vNext

# v1.4.3

- (improvement) Initialize json implementation on startup
- (improvement) Update dependencies

# v1.4.2

- (bug) Fix grpc runtime stop

# v1.4.1

- (bug) Fix grpc runtime stop

# v1.4.0

- (feature) Add `OnStopLast` to `Lifecycle`
- (improvement) Set health to unhealthy on lifecycle stop

# v1.3.2

- (feature) Add validate mehtod to custom json un/marshall functions

# v1.3.1

- (feature) Add custom json un-/marshall functions
- (internal) Use goccy/go-json when Architecture is ARM64 instead of bytedance/sonic

# v1.3.0

- (feature) Add bytedance/sonic as default fiber json en-/decoder
- (feature) Make json en-/decoder configurable with `FIBER_JSON_ENCODER` and `FIBER_JSON_DECODER`

# v1.2.0

- (feature) Specify allowed CORS headers by `CORS_ALLOW_HEADERS` as string that contains the headers list separated by `,` symbols: `Content-Type, Accept`
- (feature) Specify which headers to expose for CORS by `CORS_EXPOSE_HEADERS` as string that contains the headers list separated by `,` symbols: `Content-Type, Accept`
- (feature) Specify if CORS are allowed to contain credentials by setting the `CORS_ALLOW_CREDENTIALS` env variable to true

# v1.1.0

- (feature) Specify allowed CORS origins by `CORS_ALLOW_ORIGINS` as string that contains the hosts list separated by `,` symbols: `https://becklyn.com, https://www.becklyn.com`

# v1.0.2

- (improvement) Force fiber shutdown after 3 seconds

# v1.0.1

- (internal) Adjust logger format by `LOG_FORMAT` env variable

# v1.0.0

- (bc) Use golog instead of logrus
- (bc) Remove CQRS package

# v1.0.0-alpha.4

- (feature) Set upload limit for files by environment variable HTTP_REQUEST_BODY_LIMIT

# v1.0.0-alpha.3

- (improvement) Fiber logger uses debug level
- (improvement) Improve runtime and runtime logging

# v1.0.0-alpha.2

- (feature) Add wire sets to all packages

# v1.0.0-alpha.1

- (feature) App package
- (feature) CQRS package
- (feature) Env package
- (feature) Fiber package
- (feature) GraphQL package
- (feature) gRPC package
- (feature) Health package
- (feature) Logging package
- (feature) Metrics package
- (feature) Readyness package
