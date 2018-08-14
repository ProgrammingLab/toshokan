package main

import (
	"os"

	"github.com/ProgrammingLab/toshokan/app"
)

func main() {
	os.Exit(run())
}

func run() int {
	err := app.Run()
	if err != nil {
		return 1
	}
	return 0
}
