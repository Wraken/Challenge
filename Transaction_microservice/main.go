package main

import (
	"context"
	"log"
	"net"
	"time"

	pb_balance "Moneway_Challenge/Balance_microservice/Proto"
	pb "Moneway_Challenge/Transaction_microservice/Proto"

	"github.com/gocql/gocql"
	"google.golang.org/grpc"
)

//port of Transaction microservice
const (
	port = ":50052"
)

//Balance microservice address
const (
	grpcBalanceAddr = "localhost:50051"
)

//Client for Balance microservice
var balanceCom pb_balance.BalanceClient

//DB session
var session *gocql.Session

type server struct{}

func (s *server) MakeDeposit(ctx context.Context, in *pb.Transaction) (*pb.TransactionStatus, error) {
	//Contact Balance service to update the balance
	Ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r, err := balanceCom.DepositMoney(Ctx, &pb_balance.Transaction{AccountName: in.AccountID, NbMoney: in.Amount})
	if err != nil {
		log.Println(err)
		return &pb.TransactionStatus{Amount: 0, State: false}, nil
	}

	if r.State == false {
		return &pb.TransactionStatus{Amount: 0, State: false}, nil
	}

	//Add transtaction in db
	uuid, _ := gocql.RandomUUID()
	now := time.Now()
	err = session.Query(`INSERT INTO Transaction_service.transaction (TransactionID, AccountName, CreatedAt, Description, Amount, Notes) VALUES (?, ?, ?, ?, ?, ?)`,
		uuid, in.AccountID, now, in.Description, in.Amount, in.Notes).Exec()

	if err != nil {
		log.Println(err)
		return &pb.TransactionStatus{Amount: 0, State: false}, nil
	}

	return &pb.TransactionStatus{Amount: r.Amount, State: true}, nil
}

func (s *server) MakeCredit(ctx context.Context, in *pb.Transaction) (*pb.TransactionStatus, error) {
	//Contact Balance service to update the balance
	Ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r, err := balanceCom.CreditMoney(Ctx, &pb_balance.Transaction{AccountName: in.AccountID, NbMoney: in.Amount})
	if err != nil {
		log.Println(err)
		return &pb.TransactionStatus{Amount: 0, State: false}, nil
	}

	if r.State == false {
		return &pb.TransactionStatus{Amount: 0, State: false}, nil
	}

	//Add transtaction in db
	uuid, _ := gocql.RandomUUID()
	now := time.Now()
	err = session.Query(`INSERT INTO Transaction_service.transaction (TransactionID, AccountName, CreatedAt, Description, Amount, Notes) VALUES (?, ?, ?, ?, ?, ?)`,
		uuid, in.AccountID, now, in.Description, in.Amount, in.Notes).Exec()

	if err != nil {
		log.Println(err)
		return &pb.TransactionStatus{Amount: 0, State: false}, nil
	}

	return &pb.TransactionStatus{Amount: r.Amount, State: true}, nil
}

func (s *server) ListAllTransaction(in *pb.AllTransaction, stream pb.Transaction_ListAllTransactionServer) error {
	//Get all transaction from db
	iter := session.Query(`SELECT TransactionID,AccountName,CreatedAt,Description,Amount,Notes from Transaction_service.transaction`).Iter()

	var accName, desc, notes string
	var transactionID string
	var createdAt int64
	var amount float32

	//Send all transactions in stream
	for iter.Scan(&transactionID, &accName, &createdAt, &desc, &amount, &notes) {
		if err := stream.Send(&pb.Transaction{ID: transactionID, AccountID: accName, CreatedAt: createdAt, Description: desc, Amount: amount, Notes: notes}); err != nil {
			log.Println(err)
			return nil
		}
	}

	return nil
}

func main() {
	//Db initialisation
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.ProtoVersion = 3
	cluster.ConnectTimeout = time.Second * 20
	cluster.Consistency = gocql.One
	var err error
	session, err = cluster.CreateSession()
	if err != nil {
		log.Println(err)
		return
	}
	defer session.Close()

	//create Keyspace
	err = session.Query("CREATE KEYSPACE IF NOT EXISTS Transaction_service WITH REPLICATION = {'class' : 'SimpleStrategy', 'replication_factor' : 3};").Exec()
	if err != nil {
		log.Println(err)
		return
	}

	//create Table
	err = session.Query("CREATE TABLE IF NOT EXISTS Transaction_service.transaction (TransactionID uuid, AccountName text, CreatedAt timestamp, Description text, Amount float, Notes text, PRIMARY KEY (TransactionID, AccountName));").Exec()
	if err != nil {
		log.Println(err)
		return
	}

	// Set up a connection to the balance server.
	conn, err := grpc.Dial(grpcBalanceAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	balanceCom = pb_balance.NewBalanceClient(conn)

	//Launch grpc server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTransactionServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
