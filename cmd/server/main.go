package main

import (
	"better-shipping-app/internal/api"
	"better-shipping-app/internal/config"
	"better-shipping-app/internal/dao"
	"better-shipping-app/internal/service"
	"log"
)

func main() {
	// set up the application

	var cfg = config.LoadConfig()

	var migrate = dao.NewMigrator(cfg.DatabaseConfig)

	if err := migrate.Migrate(); err != nil {
		log.Fatal(err)
	}

	// init dao layer

	dbShell, err := dao.NewDbShell(cfg.DatabaseConfig)

	if err != nil {
		log.Fatal(err)
	}

	packDao := dao.NewPackDao(dbShell)

	// init service layer
	packService := service.NewPackService(packDao)
	shippingService := service.NewShippingService(packService)

	// init api layer
	server := api.NewServer(cfg.ServerConfig)

	api.RegisterPackApi(packService, server)
	api.RegisterShippingApi(shippingService, server)

	// start the application

	if err = server.Start(); err != nil {
		log.Fatal(err)
	}
}
