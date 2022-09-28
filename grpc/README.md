# gRPC

The `GRPC_ADDR` env variable is used to configure the address of the server. Default: `tcp://0.0.0.0:9000`.

Use the server provided by `grpc.New` as a basis for all your grpc endpoints.

```go
server := grpc.New(...)
```
