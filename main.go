package main

import (
	"grpc-startup/gateway"
	"grpc-startup/grpc/server"
	"grpc-startup/sum"
	"grpc-startup/sum/sumpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"io/ioutil"
	"net"
	"os"
)

const ADDRESS = "0.0.0.0:8080"

func main() {
	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)

	listener, err := net.Listen("tcp", ADDRESS)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	// grpc server instance creation
	svr := server.NewServer(sum.NewService())

	// service registrations
	sumpb.RegisterSumServiceServer(s, svr)

	// Serve gRPC Server
	log.Info("Serving gRPC on http://", ADDRESS)
	go func() {
		log.Fatal(s.Serve(listener))
	}()

	err = gateway.Run("dns:///" + ADDRESS)
	log.Fatalln(err)
}
