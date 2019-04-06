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
		return &pb.BalanceReply{Amount: 0, State: false}, err
	}
	bal = bal - in.NbMoney

	//Update balance in db
	err = session.Query("UPDATE balance_service.balance SET balance = ? WHERE account_name = ? and account_id = ?", bal, in.AccountName, id).Exec()
	if err != nil {
		log.Println(err)
		return &pb.BalanceReply{Amount: bal, State: false}, err
	}
	return &pb.BalanceReply{Amount: bal, State: true}, nil
}

func (s *server) DepositMoney(ctx context.Context, in *pb.Transaction) (*pb.BalanceReply, error) {
	var bal float32
	var id gocql.UUID

	//Get Balance in db
	err := session.Query("SELECT account_id,balance FROM balance_service.balance WHERE account_name = ? allow filtering", in.AccountName).Scan(&id, &bal)
	if err != nil {
		log.Println(err)
		return &pb.BalanceReply{Amount: 0, State: false}, err
	}

	bal = bal + in.NbMoney

	//Update balance in db
	err = session.Query("UPDATE balance_service.balance SET balance = ? WHERE account_name = ? and account_id = ? ", bal, in.AccountName, id).Exec()
	if err != nil {
		log.Println(err)
		return &pb.BalanceReply{Amount: bal, State: false}, nil
	}
	return &pb.BalanceReply{Amount: bal, State: true}, nil
}

func (s *server) GetBalance(ctx context.Context, in *pb.AccountName) (*pb.BalanceReply, error) {
	var bal float32

	bal = 0
	//Get Balance in db
	err := session.Query("SELECT balance FROM balance_service.balance WHERE account_name = ? allow filtering", in.AccountName).Scan(&bal)
	if err != nil {
		log.Println(err)
		return &pb.BalanceReply{Amount: 0, State: false}, err
	}

	return &pb.BalanceReply{Amount: bal, State: true}, err
}

func (s *server) CreateAccount(ctx context.Context, in *pb.AccountName) (*pb.BalanceReply, error) {
	var name string

	//Check if the account already exist in db
	err := session.Query("SELECT account_name FROM balance_service.balance WHERE account_name = ? allow filtering", in.AccountName).Scan(&name)
	if err != nil {
		log.Println(err)
	}

	if name != "" {
		return &pb.BalanceReply{Amount: 0, State: false}, err
	}

	uuid, _ := gocql.RandomUUID()
	//Create an account in db if it doesn't exist
	err = session.Query("INSERT INTO balance_service.balance(account_id,account_name,balance) VALUES(?,?,?)", uuid, in.AccountName, float32(0)).Exec()
	if err != nil {
		log.Println(err)
		return &pb.BalanceReply{Amount: 0, State: false}, err
	}

	return &pb.BalanceReply{Amount: 0, State: true}, nil
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
