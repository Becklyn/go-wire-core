//go:build wireinject

package fiber

import "github.com/google/wire"

var Default = wire.NewSet(
	wire.Struct(new(NewFiberOptions), "*"),
	wire.Struct(new(UseFiberOptions), "*"),
	NewFiber,
	NewEmptyMiddlewareHandlerMap,
)

var WithMiddleware = wire.NewSet(
	wire.Struct(new(NewFiberOptions), "*"),
	wire.Struct(new(UseFiberOptions), "*"),
	NewFiber,
)
