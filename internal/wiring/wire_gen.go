// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wiring

import (
	"base-setup/internal/app"
	"base-setup/internal/configs"
	"base-setup/internal/dataacess/database"
	"base-setup/internal/handler/http"
	"base-setup/internal/logic"
	"base-setup/internal/utils"
	"github.com/google/wire"
)

// Injectors from wire.go:

func InitializeStandaloneServer() (*app.Server, func(), error) {
	config, err := configs.NewConfig()
	if err != nil {
		return nil, nil, err
	}
	configsHTTP := config.HTTP
	checkHealthHandler := http.NewCheckHealthHandler()
	log := config.Log
	logger, cleanup, err := utils.InitializeLogger(log)
	if err != nil {
		return nil, nil, err
	}
	configsDatabase := config.Database
	db, err := database.InitializeDB(logger, configsDatabase)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	userDataAccessor := database.NewUserDataAccessor(db, logger)
	userLogic := logic.NewUserLogic(userDataAccessor, db, logger)
	userHandler := http.NewUserHandler(userLogic, logger)
	authLogic := logic.NewAuthLogic(userDataAccessor, logger)
	authHandler := http.NewAuthHandler(authLogic, logger)
	handler := http.Handler{
		CheckHealthHandler: checkHealthHandler,
		UserHandler:        userHandler,
		AuthHandler:        authHandler,
	}
	server := http.NewServer(configsHTTP, handler, logger)
	appServer := app.NewStandaloneServer(server, logger)
	return appServer, func() {
		cleanup()
	}, nil
}

// wire.go:

var WireSet = wire.NewSet(utils.WireSet, app.WireSet, configs.WireSet, http.WireSet, logic.WireSet, database.WireSet)
