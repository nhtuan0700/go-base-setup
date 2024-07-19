package configs

import "github.com/google/wire"

var WireSet = wire.NewSet(
	NewConfig,
	wire.FieldsOf(new(Config), "HTTP"),
	wire.FieldsOf(new(Config), "Log"),
	wire.FieldsOf(new(Config), "Database"),
)
