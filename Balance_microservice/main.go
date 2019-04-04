package main

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	//"github.com/scylladb/gocqlx"
	//"github.com/scylladb/gocqlx/qb"
	"github.com/gocql/gocql"
	pb "github.com/Wraken/Moneway_Challenge/Balance_microservice/Proto"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) GetBalance(ctx context.Context, in *pb.Ping) (*pb.BalanceReply, error) {
	log.Printf("Received: %v", in.Ping)
	return &pb.BalanceReply{Amount: 300}, nil
}

func main() {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.ProtoVersion = 3
	cluster.ConnectTimeout = time.Second * 4
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		log.Println(err)
		return
	}
	defer session.Close()
	log.Println("init db done")
	const personSchema = `
	CREATE TABLE IF NOT EXISTS balance (
    balance int
	)`

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
