package proto

import (
	"context"

	pb "github.com/mkaiho/go-grpc-sample2/proto"
	"google.golang.org/grpc"
)

var _ pb.TodoManagerServer = (*todoManagerServer)(nil)

func NewTodoManagerServer() *todoManagerServer {
	return &todoManagerServer{}
}

type todoManagerServer struct {
	pb.UnimplementedTodoManagerServer
}

type (
	TodoCreateRequest struct{}
)

func (s *todoManagerServer) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	return &pb.CreateResponse{
		Todo: req.Todo,
	}, nil
}

func NewTodoManagerClient(conn grpc.ClientConnInterface) pb.TodoManagerClient {
	return pb.NewTodoManagerClient(conn)
}
