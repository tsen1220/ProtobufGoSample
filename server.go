package main

import (
	calculator "grpc/test/protoc"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf(err.Error())
	}

	ser := calculator.Server{}

	grpcServer := grpc.NewServer()
	calculator.RegisterCalculatorServiceServer(grpcServer, &ser)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf(err.Error())
	}
}
