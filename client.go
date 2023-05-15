package main

import (
	"context"
	calculator "grpc/test/protoc"
	"log"

	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf(err.Error())
	}

	defer conn.Close()

	c := calculator.NewCalculatorServiceClient(conn)

	response, err := c.Sum(context.Background(), &calculator.CalculatorRequest{
		A: 101,
		B: 105,
	})

	log.Println(response)
}
