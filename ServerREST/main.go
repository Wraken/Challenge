package main

import (
	pb_Balance "Moneway_Challenge/Balance_microservice/Proto"
	pb_Transaction "Moneway_Challenge/Transaction_microservice/Proto"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

const (
	addressBalance     = "localhost:50051"
	addressTransaction = "localhost:50052"
)

type transaction struct {
	ID          string  `json:"id,omitempty"`
	AccountID   string  `json:"accountid,omitempty"`
	CreatedAt   string  `json:"createdat,omitempty"`
	Description string  `json:"description,omitempty"`
	Amount      float32 `json:"amount,omitempty"`
	Notes       string  `json:"notes,omitempty"`
}

type balance struct {
	Amount float32 `json:"amount,omitempty"`
	State  bool    `json:"State,omitempty"`
}

func createAccount(w http.ResponseWriter, r *http.Request) {
	//Pars arg from request and check if it's ok
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	if r.Form.Get("accountname") == "" {
		json.NewEncoder(w).Encode("Error: accountname is a mandatory param")
		return
	}

	//Set context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//Ask to the balance microservice to create an account
	res, err := balanceClient.CreateAccount(ctx, &pb_Balance.AccountName{AccountName: r.Form.Get("accountname")})
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode("Error: failed to create an account")
		return
	}

	if res.State == false {
		json.NewEncoder(w).Encode("Error: failed to create an account. Maybe account already exist")
		return
	}

	json.NewEncoder(w).Encode(res)
}

func getBalance(w http.ResponseWriter, r *http.Request) {
	//Set context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//Get params from request
	params := mux.Vars(r)

	//Contact Balance microservice
	res, err := balanceClient.GetBalance(ctx, &pb_Balance.AccountName{AccountName: params["id"]})
	if err != nil {
		log.Printf("Error getting balance: %v", err)
		json.NewEncoder(w).Encode("Error: failed to get balance. Maybe the account doesn't exist")
		return
	}
	if res.State == false {
		json.NewEncoder(w).Encode("Error: failed to get balance. Maybe the account doesn't exist")
		return
	}

	//Encode and send back data
	json.NewEncoder(w).Encode(&balance{res.Amount, res.State})
}

func getAllTransactions(w http.ResponseWriter, r *http.Request) {
	//Set context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//Contact Transaction microservice
	stream, err := transactionClient.ListAllTransaction(ctx, &pb_Transaction.AllTransaction{})
	if err != nil {
		log.Println(err)
		return
	}

	transactions := make([]transaction, 0)

	//Read transaction from stream
	for {
		r, err := stream.Recv()
		if err == io.EOF {
			break
		}
		time := time.Unix(r.CreatedAt, 0)
		transactions = append(transactions, transaction{ID: r.ID, AccountID: r.AccountID, CreatedAt: time.String(),
			Description: r.Description, Amount: r.Amount, Notes: r.Notes})
	}
	json.NewEncoder(w).Encode(transactions)
}

func debitAccount(w http.ResponseWriter, r *http.Request) {
	//Pars arg from request
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	amount, err := strconv.ParseFloat(r.Form.Get("amount"), 32)
	if err != nil {
		log.Println(err)
		return
	}

	//Set context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if r.Form.Get("accountid") == "" {
		json.NewEncoder(w).Encode("Error: accountid is mandatory")
		return
	}

	//Contact transaction microservice to credit the account
	res, err := transactionClient.MakeCredit(ctx, &pb_Transaction.Transaction{
		ID: "", AccountID: r.Form.Get("accountid"), CreatedAt: 0, Amount: float32(amount), Notes: r.Form.Get("notes")})
	if err != nil {
		log.Printf("Could not make transaction: %v", err)
		json.NewEncoder(w).Encode("Error: deposit failed, checks your params")
		return
	}

	if res.State == false {
		json.NewEncoder(w).Encode("Error: deposit failed, checks your params")
		return
	}

	json.NewEncoder(w).Encode(&balance{res.Amount, res.State})
}

func creditAccount(w http.ResponseWriter, r *http.Request) {
	//Pars arg
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	amount, err := strconv.ParseFloat(r.Form.Get("amount"), 32)
	if err != nil {
		log.Println(err)
		return
	}

	//Set context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if r.Form.Get("accountid") == "" {
		json.NewEncoder(w).Encode("Error: accountid is mandatory")
		return
	}
	//Contact transaction microservice to make a deposit on the account
	res, err := transactionClient.MakeDeposit(ctx, &pb_Transaction.Transaction{
		ID: "", AccountID: r.Form.Get("accountid"), CreatedAt: 0, Amount: float32(amount), Notes: r.Form.Get("notes")})
	if err != nil {
		log.Printf("Could not make transaction: %v", err)
		json.NewEncoder(w).Encode("Error: deposit failed, checks your params")
		return
	}

	if res.State == false {
		json.NewEncoder(w).Encode("Error: deposit failed, checks your params")
		return
	}

	json.NewEncoder(w).Encode(&balance{res.Amount, res.State})
}

//Client for Balance microservice
var balanceClient pb_Balance.BalanceClient

//Client for transaction
var transactionClient pb_Transaction.TransactionClient

func main() {
	// Set up a connection to the balance server.
	conn, err := grpc.Dial(addressBalance, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	balanceClient = pb_Balance.NewBalanceClient(conn)

	//Set up connection to the transaction server
	conn2, err := grpc.Dial(addressTransaction, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Transaction connect fail: %v", err)
	}
	defer conn2.Close()
	transactionClient = pb_Transaction.NewTransactionClient(conn2)

	router := mux.NewRouter()
	router.HandleFunc("/balance/{id}", getBalance).Methods("GET")
	router.HandleFunc("/balance/createaccount", createAccount).Methods("POST")
	router.HandleFunc("/transactions", getAllTransactions).Methods("GET")
	router.HandleFunc("/transactions/debitaccount", debitAccount).Methods("POST")
	router.HandleFunc("/transactions/creditaccount", creditAccount).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}
