package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/NoNamePL/GoWalletExchanger/api/gw-wallet-exchanger"
	"github.com/NoNamePL/GoWalletExchanger/iternal/handlers/wallet/walhandlers"
	"github.com/NoNamePL/GoWalletExchanger/iternal/middleware/logger"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
	"google.golang.org/grpc"
)

func main() {
	// start handler gin server
	// router := gin.Default()
	router := gin.New()

	// create logger
	logger, err := logger.InitLogger("REST-API")
	if err != nil {
		log.Fatal("can't open/create/append to log file")
	}

	// create config file
	// cfg, err := config.NewConfig()
	// if err != nil {
	// 	logger.Error(err.Error())
	// 	os.Exit(1)
	// }

	router.Use(sloggin.New(logger))
	router.Use(gin.Recovery())

	// connect to DB
	// db, err := postgres.ConnectDB(cfg)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	test := sql.DB{}

	db := &test

	// start grpc client
	conn, err := grpc.Dial("localhost:50001", grpc.WithInsecure())
	if err != nil {
		logger.Error("didn't connect to rpc: %v", err)
		os.Exit(1)
	}
	defer conn.Close()

	grpcClient := pb.NewExchangeServiceClient(conn)

	// пример вызова GetExchangeRates
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	rates, err := grpcClient.GetExchangeRates(ctx, &pb.Empty{})
	if err != nil {
		logger.Error("couldn't get rates:", "error", err)
	}
	logger.Info("Exchange Rates: ", "string", rates.Rates)

	// Пример вызова GetExchangeRateForCurrency
	currencyReq := &pb.CurrencyRequest{FromCurrency: "USD", ToCurrency: "EUR"}
	ratesResp, err := grpcClient.GetExchangeRateForCurrency(ctx, currencyReq)
	if err != nil {
		logger.Error("couldn't get exchange rate:", "error", err)
	}
	logger.Info(fmt.Sprintf("Exchange Rate from %s to %s: %f", ratesResp.FromCurrency, ratesResp.ToCurrency, ratesResp.Rate))

	walhandlers.RegisterHandlers(router, db, &grpcClient, logger)

	router.Run()

	// // create grpc server
	// grpcServer := grpc.NewServer()

	// reflection.Register(grpcServer)

	// // create grpc client
	// // grpc.NewClient()

	// // pb.NewExchangeServiceClient(grpcServer, &ClientServerObject)

	// client := pb.NewExchangeServiceClient(&grpc.ClientConn{})

	// grpcServer.

	// pb.ExchangeServiceClient{}

}

// type ClientServerObject struct {
// 	pb.ExchangeServiceClient
// }

// func (cl *ClientServerObject) GetExchangeRates(ctx context.Context, in *pb.Empty) (*pb.ExchangeRatesResponse, error) {
// 	return &pb.ExchangeRatesResponse{
// 		Rates: map[string]float32{"test": 21.3},
// 	}, nil
// }

// func (cl *ClientServerObject) GetExchangeRateForCurrency(ctx context.Context, in *pb.CurrencyRequest) (*pb.ExchangeRateResponse, error) {
// 	return &pb.ExchangeRateResponse{
// 		FromCurrency: "EURO",
// 		ToCurrency:   "KZT",
// 		Rate:         21.4,
// 	}, nil
// }
