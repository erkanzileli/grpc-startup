package main

import (
	"grpc-startup/sum/sumpb"
	"context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	cc, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer cc.Close()

	//sum
	sumClient := sumpb.NewSumServiceClient(cc)
	sumReq := &sumpb.SumRequest{
		Number1: 3,
		Number2: 2,
	}
	sumRes, err := sumClient.Sum(context.TODO(), sumReq)
	if err != nil {
		log.Fatalln("failed to sum", err)
	}
	log.Println("Sum Response:", sumRes.Result)
}
