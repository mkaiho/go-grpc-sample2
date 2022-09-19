package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/mkaiho/go-grpc-sample2/internal/adapter/proto"
	pb "github.com/mkaiho/go-grpc-sample2/proto"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 3000, "listening port")
)

func init() {
	flag.Parse()
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterTodoManagerServer(grpcServer, proto.NewTodoManagerServer())
	grpcServer.Serve(lis)
	log.Printf("listening port: %d", *port)
}
