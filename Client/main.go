package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	pb_Balance "github.com/Wraken/Moneway_Challenge/Balance_microservice/Proto"
)

const (
	address     = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb_Balance.NewBalanceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetBalance(ctx, &pb_Balance.Ping{Ping: "test"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Get Balance response: %d", r.Amount)
}
