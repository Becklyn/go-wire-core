//go:build wireinject

package readyness

import "github.com/google/wire"

var Default = wire.NewSet(
	New,
	wire.Struct(new(UseFiberEndpointOptions), "*"),
)
