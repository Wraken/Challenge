package main

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	//"github.com/scylladb/gocqlx"
	//"github.com/scylladb/gocqlx/qb"
	pb "Moneway_Challenge/Balance_microservice/Proto"

	"github.com/gocql/gocql"
)

//microservice port
const (
	port = ":50051"
)

var session *gocql.Session

type server struct{}

func (s *server) CreditMoney(ctx context.Context, in *pb.Transaction) (*pb.BalanceReply, error) {

	var bal float32
	var id gocql.UUID

	//get balance in db
	err := session.Query("SELECT account_id,balance FROM balance_service.balance WHERE account_name = ? allow filtering", in.AccountName).Scan(&id, &bal)
	if err != nil {
		log.Println(err)
	}
	bal = bal - in.NbMoney

	//Update balance in db
	err = session.Query("UPDATE balance_service.balance SET balance = ? WHERE account_name = ? and account_id = ?", bal, in.AccountName, id).Exec()
	if err != nil {
		log.Println(err)
	}
	log.Println("Credit")
	return &pb.BalanceReply{Amount: bal}, nil
}

func (s *server) DepositMoney(ctx context.Context, in *pb.Transaction) (*pb.BalanceReply, error) {
	var bal float32
	var id gocql.UUID

	//Get Balance in db
	err := session.Query("SELECT account_id,balance FROM balance_service.balance WHERE account_name = ? allow filtering", in.AccountName).Scan(&id, &bal)
	if err != nil {
		log.Println(err)
	}
	bal = bal + in.NbMoney

	//Update balance in db
	err = session.Query("UPDATE balance_service.balance SET balance = ? WHERE account_name = ? and account_id = ? ", bal, in.AccountName, id).Exec()
	if err != nil {
		log.Println(err)
	}
	log.Println("Deposit")
	return &pb.BalanceReply{Amount: bal}, nil
}

func (s *server) GetBalance(ctx context.Context, in *pb.AccountName) (*pb.BalanceReply, error) {
	var bal float32

	bal = 0
	err := session.Query("SELECT balance FROM balance_service.balance WHERE account_name = ? allow filtering", in.AccountName).Scan(&bal)
	if err != nil {
		log.Println(err)
	}

	return &pb.BalanceReply{Amount: bal}, err
}

func main() {
	//db initialisation
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
	log.Println("init db done")

	//create Keyspace
	err = session.Query("CREATE KEYSPACE IF NOT EXISTS Balance_service WITH REPLICATION = {'class' : 'SimpleStrategy', 'replication_factor' : 3};").Exec()
	if err != nil {
		log.Println(err)
		return
	}

	//create Table
	err = session.Query("CREATE TABLE IF NOT EXISTS Balance_service.balance (account_id uuid, account_name text, balance float, PRIMARY KEY (account_id, account_name));").Exec()
	if err != nil {
		log.Println(err)
		return
	}

	//Start server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterBalanceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
