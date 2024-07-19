package handler

import "github.com/google/wire"

var WireSet = wire.NewSet(
	NewCheckHealthHandler,
	NewUserHandler,
	NewPostHandler,
	wire.Struct(new(Handler), "*"),
)
