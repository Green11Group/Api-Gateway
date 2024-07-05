package main

import (
	"apigateway/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conngar, err := grpc.NewClient(":50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conngar.Close()

	conncom, err := grpc.NewClient(":50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conncom.Close()

	connuser, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer connuser.Close()

	connsus, err := grpc.NewClient(":50055", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer connuser.Close()

	router:=api.NewRouter(conngar,conncom,connuser,connsus)
	router.Run(":8080")
}
