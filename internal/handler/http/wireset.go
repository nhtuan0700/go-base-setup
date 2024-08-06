package http

import "github.com/google/wire"

var WireSet = wire.NewSet(
	NewServer,
	NewCheckHealthHandler,
	NewUserHandler,
	// NewPostHandler,
	wire.Struct(new(Handler), "*"),
)
