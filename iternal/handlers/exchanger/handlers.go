package exhandler

import GoWalletExchanger "github.com/NoNamePL/GoWalletExchanger/api/gw-wallet-exchanger"

type Server struct {
	GoWalletExchanger.UnimplementedExchangeServiceServer
}

// func (s *Server) 