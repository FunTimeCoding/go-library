package coder

import (
	"context"
	"errors"
	errorLibrary "github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/monitor/coder/server"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Run() error {
	if len(os.Args) < 2 {
		return errors.New("listen address missing")
	}

	l, listenFail := net.Listen("tcp", os.Args[1])
	errorLibrary.PanicOnError(listenFail)
	log.Printf("listen on ws://%v", l.Addr())
	s := &http.Server{
		Handler:      server.Server{Logf: log.Printf},
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}
	channelFail := make(chan error, 1)

	go func() {
		log.Println("serve")
		channelFail <- s.Serve(l)
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	log.Println("wait for signal")

	select {
	case selectFail := <-channelFail:
		log.Printf("serve fail: %v", selectFail)
	case sig := <-signals:
		log.Printf("terminate: %v", sig)
	}

	c, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	return s.Shutdown(c)
}
