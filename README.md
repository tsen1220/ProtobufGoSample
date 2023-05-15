# Protobuf Sample With GO

## Before starting

We need to get some gRPC package

```
	go get github.com/golang/protobuf v1.5.3
	go get google.golang.org/grpc v1.55.0
```

## Introduction

Write protobuf file, following proto3 documents: https://protobuf.dev/programming-guides/proto3/
For example with `test.proto`: 

```
syntax = "proto3";  // protocol buffer version

package calculator;  // for name space
option go_package = "calculator";  // generated Go code
option csharp_namespace = "calculator";  // generated C# code

message CalculatorRequest {
  int64 a = 1;
  int64 b = 2;
}

message CalculatorResponse {
  int64 result = 1;
}

service CalculatorService {
  rpc Sum(CalculatorRequest) returns (CalculatorResponse) {};
}
```

Build protobuf message with .proto file

```
protoc ./protoc/*.proto --go_out=plugins=grpc:.
```

Then implement the pb interface. 
For example, create test.go file and implement interface with `Server` struct.
The interface is from generated test.pb.go file.

```
type Server struct{}

func (s *Server) Sum(ctx context.Context, req *CalculatorRequest) (*CalculatorResponse, error) {
	return &CalculatorResponse{Result: req.A + req.B}, nil
}
```

After implementing, create `server.go` and start gRPC server. Registering service on gRPC server.

```
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf(err.Error())
	}

	ser := calculator.Server{} // implemented protobuf interface

	grpcServer := grpc.NewServer()
	calculator.RegisterCalculatorServiceServer(grpcServer, &ser) // registering

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf(err.Error())
	}
```

Then we can test with `client.go`. Connected with gRPC server 

```
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

	log.Println(response.Result)
```

We will get response

```
206
```