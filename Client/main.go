package main

import (
	"context"
	"io"
	"log"
	"time"

	pb_Balance "Moneway_Challenge/Balance_microservice/Proto"
	pb_Transaction "Moneway_Challenge/Transaction_microservice/Proto"

	"google.golang.org/grpc"
)

const (
	addressBalance     = "localhost:50051"
	addressTransaction = "localhost:50052"
)

func main() {
	// Set up a connection to the balance server.
	conn, err := grpc.Dial(addressBalance, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	balanceClient := pb_Balance.NewBalanceClient(conn)

	// Contact the server and print out its response.
	ctxBalance, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn2, err := grpc.Dial(addressTransaction, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Transaction connect fail: %v", err)
	}
	defer conn2.Close()
	transactionClient := pb_Transaction.NewTransactionClient(conn2)

	r, err1 := balanceClient.GetBalance(ctxBalance, &pb_Balance.AccountName{AccountName: "Test"})
	if err1 != nil {
		log.Printf("could not: %v", err1)
	}
	log.Printf("Get Balance response: %f", r.Amount)

	res, err2 := transactionClient.MakeDeposit(ctxBalance, &pb_Transaction.Transaction{
		ID: "", AccountID: "Test", CreatedAt: 0, Description: "description du test", Amount: 10, Notes: "ceci est un test"})
	if err2 != nil {
		log.Printf("Could not make transaction: %v", err2)
	}

	log.Printf("Balance after Deposit: %f and status %t", res.Amount, res.State)

	res, err4 := transactionClient.MakeCredit(ctxBalance, &pb_Transaction.Transaction{
		ID: "", AccountID: "Test", CreatedAt: 0, Description: "Credit test", Amount: 10, Notes: "ceci est un test"})
	if err4 != nil {
		log.Fatalf("Could not make transaction: %v", err4)
	}

	log.Printf("Balance after Deposit: %f and status %t", res.Amount, res.State)

	stream, err3 := transactionClient.ListAllTransaction(ctxBalance, &pb_Transaction.AllTransaction{})
	if err3 != nil {
		log.Println(err3)
		return
	}
	for {
		transaction, err := stream.Recv()
		if err == io.EOF {
			break
		}
		log.Printf("%s, %s, %s, %f", transaction.ID, transaction.AccountID, transaction.Description, transaction.Amount)
	}

}
