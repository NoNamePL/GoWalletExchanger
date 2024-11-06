package main

import (
	"log"

	"github.com/NoNamePL/GoWalletExchanger/iternal/config"
	"github.com/NoNamePL/GoWalletExchanger/iternal/handlers"
	"github.com/NoNamePL/GoWalletExchanger/iternal/storages/postgres"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.ConnectDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	handlers.RegisterHandlers(router, db)

	router.Run()
}
