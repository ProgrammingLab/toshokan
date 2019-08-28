package server

import (
	"context"

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

func (s *lendingServiceServerImpl) ListUserLendings(_ context.Context, _ *api_pb.ListUserLendingsRequest) (*api_pb.ListLendingsResponse, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *lendingServiceServerImpl) GetLending(_ context.Context, _ *api_pb.GetLendingRequest) (*api_pb.Lending, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *lendingServiceServerImpl) BorrowBook(_ context.Context, _ *api_pb.BorrowBookRequest) (*api_pb.Lending, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *lendingServiceServerImpl) ReturnBook(_ context.Context, _ *api_pb.ReturnBookRequest) (*api_pb.Lending, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}
