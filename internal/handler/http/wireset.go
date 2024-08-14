package http

import "github.com/google/wire"

var WireSet = wire.NewSet(
	NewServer,
	NewCheckHealthHandler,
	NewUserHandler,
	NewAuthHandler,
	wire.Struct(new(Handler), "*"),
)
