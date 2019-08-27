package main

import (
	"os"

	"google.golang.org/grpc/grpclog"

	"github.com/ProgrammingLab/toshokan/app"
)

func main() {
	os.Exit(run())
}

func run() int {
	err := app.Run()
	if err != nil {
		grpclog.Errorf("server was shutdown with errors: %v", err)
		return 1
	}
	return 0
}
