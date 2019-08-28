package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "github.com/ProgrammingLab/toshokan/api"
)

// NewLendingServiceServer creates a new LendingServiceServer instance.
func NewLendingServiceServer() interface {
	api_pb.LendingServiceServer
	grapiserver.Server
} {
	return &lendingServiceServerImpl{}
}

type lendingServiceServerImpl struct {
}

func (s *lendingServiceServerImpl) ListLendings(ctx context.Context, req *api_pb.ListLendingsRequest) (*api_pb.ListLendingsResponse, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *lendingServiceServerImpl) GetLending(ctx context.Context, req *api_pb.GetLendingRequest) (*api_pb.Lending, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *lendingServiceServerImpl) CreateLending(ctx context.Context, req *api_pb.CreateLendingRequest) (*api_pb.Lending, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *lendingServiceServerImpl) UpdateLending(ctx context.Context, req *api_pb.UpdateLendingRequest) (*api_pb.Lending, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *lendingServiceServerImpl) DeleteLending(ctx context.Context, req *api_pb.DeleteLendingRequest) (*empty.Empty, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}
