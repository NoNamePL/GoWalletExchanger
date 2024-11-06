package main

import (
	"log"

	"github.com/NoNamePL/GoWalletExchanger/iternal/config"
	storage "github.com/NoNamePL/GoWalletExchanger/iternal/storages"
	"github.com/NoNamePL/GoWalletExchanger/iternal/storages/postgres"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}


	db, err := postgres.ConnectDB(cfg)

	conn := grpc.NewServer()

}