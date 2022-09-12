package main

import (
	"info/pkg/api"
	"info/pkg/api/grpc"
)

func main() {
	opt := options()
	serviceProvider := api.RegisterServices(opt)

	grpcServer := grpc.NewServer(serviceProvider)
	grpcServer.Run()
}
