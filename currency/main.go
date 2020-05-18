package main

import (
	hclog "github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	protos "microsrv/currency/protos/currency"
	"microsrv/currency/server"
	"net"
	"os"
)

func main() {
	l := hclog.Default()

	gs := grpc.NewServer()
	cs := server.NewCurrency(l)

	protos.RegisterCurrencyServer(gs, cs)

	reflection.Register(gs)

	log, err := net.Listen("tcp", ":9092")
	if err != nil {
		l.Error("[ERROR]", err)
		os.Exit(1)
	}
	gs.Serve(log)
}
