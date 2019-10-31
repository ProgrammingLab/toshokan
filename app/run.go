package app

import (
	"github.com/ProgrammingLab/toshokan/app/server"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"github.com/srvc/appctx"
)

// Run starts the grapiserver.
func Run() error {
	ctx := appctx.Global()

	s := grapiserver.New(
		grapiserver.WithDefaultLogger(),
		grapiserver.WithServers(
			server.NewBookServiceServer(),
			server.NewLendingServiceServer(),
		),
	)
	return s.Serve(ctx)
}
