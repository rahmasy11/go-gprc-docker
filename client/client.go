package main

import (
    "context"
    "log"
    "time"

    "go-grpc-docker/proto"
    "google.golang.org/grpc"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }
    defer conn.Close()

    client := proto.NewGreetServiceClient(conn)
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    response, err := client.SayHello(ctx, &proto.GreetRequest{Name: "World"})
    if err != nil {
        log.Fatalf("Failed to greet: %v", err)
    }

    log.Printf("Response: %s", response.Message)
}
