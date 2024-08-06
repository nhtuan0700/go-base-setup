//go:build wireinject
// +build wireinject

//
//go:generate go run github.com/google/wire/cmd/wire
package wiring

import (
	"base-setup/internal/app"
	"base-setup/internal/configs"
	"base-setup/internal/dataacess/database"

	"base-setup/internal/utils"
	"base-setup/internal/handler/http"
	"base-setup/internal/logic"

	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	utils.WireSet,
	app.WireSet,
	configs.WireSet,
	http.WireSet,
	logic.WireSet,
	database.WireSet,
)

func InitializeStandaloneServer() (*app.Server, func(), error) {
	wire.Build(WireSet)

	return nil, nil, nil
}
