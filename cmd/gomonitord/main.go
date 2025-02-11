package main

import (
	errorLibrary "github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/monitor/coder"
	"github.com/funtimecoding/go-library/pkg/monitor/gorilla"
	"log"
)

func main() {
	if false {
		// Start server: go run cmd/gomonitord/main.go localhost:3002
		// Connect: websocat ws://127.0.0.1:3002 --protocol echo
		log.SetFlags(0)
		errorLibrary.PanicOnError(coder.Run())
	}

	if true {
		gorilla.Run()
	}
}
