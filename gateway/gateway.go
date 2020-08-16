package gateway

import (
	"grpc-startup/sum/sumpb"
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"io/ioutil"
	"net/http"
	"os"
)

const ADDRESS = "0.0.0.0:8081"

func Run(dialAddr string) error {
	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)

	opts := []grpc.DialOption{grpc.WithInsecure(), grpc.WithBlock()}
	conn, err := grpc.DialContext(
		context.Background(),
		dialAddr,
		opts...,
	)
	if err != nil {
		return fmt.Errorf("failed to dial server: %w", err)
	}

	mux := runtime.NewServeMux()
	err = sumpb.RegisterSumServiceHandler(context.Background(), mux, conn)
	if err != nil {
		return fmt.Errorf("failed to register gateway: %w", err)
	}

	gwServer := &http.Server{
		Addr: ADDRESS,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			mux.ServeHTTP(w, r)
		}),
	}

	return fmt.Errorf("serving gRPC-Gateway server: %w", gwServer.ListenAndServe())
}
