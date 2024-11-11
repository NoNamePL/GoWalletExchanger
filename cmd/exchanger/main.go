package main

import (
	"context"
	"log"
	"net"

	pb "github.com/NoNamePL/GoWalletExchanger/api/gw-wallet-exchanger"
	"github.com/NoNamePL/GoWalletExchanger/iternal/config"
	"github.com/NoNamePL/GoWalletExchanger/iternal/storages/postgres"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type ExchangeService struct {
	pb.UnimplementedExchangeServiceServer
}

func (ex *ExchangeService) GetExchangeRates(ctx context.Context, in *pb.Empty) (*pb.ExchangeRatesResponse, error) {
	rates := map[string]float32{
		"USD": 1.0,
		"EUR": 0.85,
		"JPY": 110.0,
	}

	return &pb.ExchangeRatesResponse{Rates: rates}, nil
}

func (ex *ExchangeService) GetExchangeRateForCurrency(ctx context.Context, req *pb.CurrencyRequest) (*pb.ExchangeRateResponse, error) {
	rates := map[string]float32{
		"USD": 1.0,
		"EUR": 0.85,
		"JPY": 110.0,
	}

	fromRate, fromExists := rates[req.FromCurrency]
	toRate, toExists := rates[req.ToCurrency]

	if !fromExists || !toExists {
		return nil, status.Errorf(codes.NotFound, "Currency not found")

	}

	exchangeRate := toRate / fromRate

	return &pb.ExchangeRateResponse{
		FromCurrency: req.FromCurrency,
		ToCurrency:   req.ToCurrency,
		Rate:         float32(exchangeRate),
	}, nil
}

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.ConnectDB(cfg)

	// create grpc server on 9000 port
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	pb.RegisterExchangeServiceServer(grpcServer, &ExchangeService{})

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}

	// conn,err :=grpc.NewClient("localhost:50051",grpc.WithInsecure())
	// if err !=nil{
	// 	log.Fatal(err.Error())
	// }
	// defer conn.Close()

	// s := grpc.NewServer()

	// pb.RegisterExchangeServiceServer(s,&)

	// client := pb.NewExchangeServiceClient(conn)
}
