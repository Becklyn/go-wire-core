//go:build wireinject

package graphql

import "github.com/google/wire"

var Default = wire.NewSet(
	NewQuery,
	NewMutation,
	NewSubscribtion,
	wire.Struct(new(NewSchemaOptions), "*"),
	NewSchema,
)
