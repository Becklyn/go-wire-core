//go:build wireinject

package health

import "github.com/google/wire"

var Default = wire.NewSet(
	New,
	wire.Struct(new(UseFiberEndpointOptions), "*"),
)
