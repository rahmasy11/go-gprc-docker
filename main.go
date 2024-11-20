package main

import (
    "context"
    "log"
    "net"

    "go-grpc-docker/proto"
    "google.golang.org/grpc"
)

type server struct {
    proto.UnimplementedGreetServiceServer
}

func (s *server) SayHello(ctx context.Context, req *proto.GreetRequest) (*proto.GreetResponse, error) {
    log.Printf("Received request: %s", req.Name)
    return &proto.GreetResponse{Message: "Hello, " + req.Name}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    s := grpc.NewServer()
    proto.RegisterGreetServiceServer(s, &server{})

    log.Println("Server is running on port 50051")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
