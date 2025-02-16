package main

import (
	errorLibrary "github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/monitor/coder"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/gorilla"
	"github.com/funtimecoding/go-library/pkg/monitor/gorilla/example_client"
)

func main() {
	if false {
		// Start server: go run cmd/gomonitord/main.go localhost:3002
		// Connect: websocat ws://127.0.0.1:3002 --protocol echo
		errorLibrary.PanicOnError(coder.Run())
	}

	if false {
		example_client.Run()
	}

	if true {
		gorilla.Run(constant.Address)
	}
}
