package database

import "github.com/google/wire"

var WireSet = wire.NewSet(InitializeDB, NewUserDataAccessor)
