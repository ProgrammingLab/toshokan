package app

import (
	"github.com/ProgrammingLab/toshokan/app/server"
	"github.com/izumin5210/grapi/pkg/grapiserver"
)

// Run starts the grapiserver.
func Run() error {
	s := grapiserver.New(
		grapiserver.WithDefaultLogger(),
		grapiserver.WithServers(
			server.NewUserServiceServer(),
			server.NewSessionServiceServer(),
		),
	)
	return s.Serve()
}
