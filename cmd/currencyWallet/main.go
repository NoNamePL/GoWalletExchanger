package main

import (
	"context"
	"log"
	"time"

	pb "github.com/NoNamePL/GoWalletExchanger/api/gw-wallet-exchanger"
	"github.com/NoNamePL/GoWalletExchanger/iternal/config"
	walhandlers "github.com/NoNamePL/GoWalletExchanger/iternal/handlers/wallet"
	"github.com/NoNamePL/GoWalletExchanger/iternal/storages/postgres"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// start handler gin server
	router := gin.Default()

	// create config
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	// connect to DB
	db, err := postgres.ConnectDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// start grpc client
	conn, err := grpc.Dial("localhost:50001",grpc.WithInsecure())
	if err != nil {
		log.Fatalf("didn't connect to rpc: %v",err)
	}
	defer conn.Close()

	grpcClient := pb.NewExchangeServiceClient(conn)


	// пример вызова GetExchangeRates
	ctx,cancel := context.WithTimeout(context.Background(),time.Second)
	defer cancel()
	
	rates,err := grpcClient.GetExchangeRates(ctx,&pb.Empty{})
	if err != nil {
		log.Fatalf("couldn't get rates: %v",err)
	}
	log.Printf("Exchange Rates: %v", rates.Rates)

	// Пример вызова GetExchangeRateForCurrency
	currencyReq := &pb.CurrencyRequest{FromCurrency: "USD", ToCurrency: "EUR"}
	ratesResp,err := grpcClient.GetExchangeRateForCurrency(ctx,currencyReq)
	if err != nil {
		log.Fatalf("couldn't get exchange rate: %v", err)
	}
	log.Println("Exchange Rate from %s to %s: %f",ratesResp.FromCurrency,ratesResp.ToCurrency,ratesResp.Rate)
	

	walhandlers.RegisterHandlers(router, db,&grpcClient)

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

type ClientServerObject struct {
	pb.ExchangeServiceClient
}

func (cl *ClientServerObject) GetExchangeRates(ctx context.Context, in *pb.Empty) (*pb.ExchangeRatesResponse, error) {
	return &pb.ExchangeRatesResponse{
		Rates: map[string]float32{"test": 21.3},
	}, nil
}

func (cl *ClientServerObject) GetExchangeRateForCurrency(ctx context.Context, in *pb.CurrencyRequest) (*pb.ExchangeRateResponse, error) {
	return &pb.ExchangeRateResponse{
		FromCurrency: "EURO",
		ToCurrency:   "KZT",
		Rate:         21.4,
	}, nil
}
