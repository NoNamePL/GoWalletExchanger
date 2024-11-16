package main

import (
	"context"
	"database/sql"
	"net"
	"os"

	pb "github.com/NoNamePL/GoWalletExchanger/api/gw-wallet-exchanger"
	"github.com/NoNamePL/GoWalletExchanger/iternal/middleware/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type ExchangeService struct {
	pb.UnimplementedExchangeServiceServer
	db *sql.DB
}

func (ex *ExchangeService) GetExchangeRates(ctx context.Context, in *pb.Empty) (*pb.ExchangeRatesResponse, error) {
	rates := map[string]float32{
		"USD": 1.0,
		"EUR": 0.85,
		"JPY": 110.0,
	}

	// ex.db.Query()

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

	// create logger
	logger, err := logger.InitLogger("GRPC")
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	// create config file
	// cfg, err := config.NewConfig()
	// if err != nil {
	// 	logger.Error(err.Error())
	// 	os.Exit(1)
	//

	// connect to DB
	// db, err := postgres.ConnectDB(cfg)
	// if err != nil {
	// 	logger.Error("can't connect to db")
	// 	os.Exit(1)
	// }

	test := sql.DB{}

	db := &test

	// create grpc server on 9000 port
	lis, err := net.Listen("tcp", ":50001")
	if err != nil {
		logger.Error("can't start grpc server")
		os.Exit(1)
	}

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	pb.RegisterExchangeServiceServer(grpcServer, &ExchangeService{db: db})

	err = grpcServer.Serve(lis)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
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
