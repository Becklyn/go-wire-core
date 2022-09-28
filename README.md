# golang wire core package

## Installation

Adding _go-wire-core_ to a Go project is as easy as calling this command

```shell
go get -u github.com/Becklyn/go-wire-core
```

## Packages

We provide packages for common use cases:

- `cqrs`: Basic interfaces for apps that do cqrs
- `env`: Environment variables processing
- `fiber`: Webserver
- `graphql`: GraphQL
- `grpc`: Basis of a gRPC server
- `health`: Health indication
- `logging`: Logging
- `metrics`: Automatic webserver metrics
- `readyness`: Readyness indication

You can see detailed documentation in the `README.md` based at the root of the packages.

## Usage

We recommend using the `app.Runtime` in your application. The app documentation gives an example of how to integrate the runtime and other packages into your app: [Documentation](app/README.md)

## Integrated 3rd party libraries

The list of 3rd party libraries that we provide as packages by this package:

- Fiber webserver - https://github.com/gofiber/fiber
- GraphQL - https://github.com/graphql-go/graphql
- gRPC - https://github.com/grpc/grpc-go
- Logrus - https://github.com/sirupsen/logrus
- Prometheus metrics - https://github.com/prometheus/client_golang
