package server

import (
	"context"
	"testing"

	"github.com/golang/protobuf/ptypes/empty"

	api_pb "github.com/ProgrammingLab/toshokan/api"
)

func Test_LendingServiceServer_ListLendings(t *testing.T) {
	svr := NewLendingServiceServer()

	ctx := context.Background()
	req := &api_pb.ListLendingsRequest{}

	resp, err := svr.ListLendings(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_LendingServiceServer_GetLending(t *testing.T) {
	svr := NewLendingServiceServer()

	ctx := context.Background()
	req := &api_pb.GetLendingRequest{}

	resp, err := svr.GetLending(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_LendingServiceServer_CreateLending(t *testing.T) {
	svr := NewLendingServiceServer()

	ctx := context.Background()
	req := &api_pb.CreateLendingRequest{}

	resp, err := svr.CreateLending(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_LendingServiceServer_UpdateLending(t *testing.T) {
	svr := NewLendingServiceServer()

	ctx := context.Background()
	req := &api_pb.UpdateLendingRequest{}

	resp, err := svr.UpdateLending(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_LendingServiceServer_DeleteLending(t *testing.T) {
	svr := NewLendingServiceServer()

	ctx := context.Background()
	req := &api_pb.DeleteLendingRequest{}

	resp, err := svr.DeleteLending(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}
