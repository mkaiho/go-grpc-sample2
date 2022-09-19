package main

import (
	"context"
	"flag"
	"log"

	"github.com/mkaiho/go-grpc-sample2/internal/adapter/proto"
	pbmodel "github.com/mkaiho/go-grpc-sample2/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:3000", "server address")
)

func init() {
	flag.Parse()
}

func main() {
	ctx := context.Background()
	log.Printf("server address: %s", *addr)
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(*addr, opts...)
	if err != nil {
		log.Fatal(err)
	}
	client := proto.NewTodoManagerClient(conn)
	defer conn.Close()

	resp, err := client.Create(ctx, &pbmodel.CreateRequest{
		Todo: &pbmodel.Todo{
			Name: "work",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("resp: %v", resp)
}
