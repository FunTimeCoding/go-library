package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/coder/websocket"
	"golang.org/x/time/rate"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// echoServer is the WebSocket echo server implementation.
// It ensures the client speaks the echo subprotocol and
// only allows one message every 100ms with a 10 message burst.
type echoServer struct {
	logf func(
		f string,
		v ...any,
	)
}

func (s echoServer) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	c, err := websocket.Accept(
		w,
		r,
		&websocket.AcceptOptions{Subprotocols: []string{"echo"}},
	)

	if err != nil {
		s.logf("%v", err)

		return
	}

	defer c.CloseNow()

	if c.Subprotocol() != "echo" {
		c.Close(
			websocket.StatusPolicyViolation,
			"client must speak the echo subprotocol",
		)

		return
	}

	l := rate.NewLimiter(rate.Every(time.Millisecond*100), 10)

	for {
		err = echo(c, l)

		if websocket.CloseStatus(err) == websocket.StatusNormalClosure {
			return
		}

		if err != nil {
			s.logf("failed to echo with %v: %v", r.RemoteAddr, err)

			return
		}
	}
}

// echo reads from the WebSocket connection and then writes
// the received message back to it.
// The entire function has 10s to complete.
func echo(
	c *websocket.Conn,
	l *rate.Limiter,
) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	err := l.Wait(ctx)

	if err != nil {
		return err
	}

	typ, r, err := c.Reader(ctx)

	if err != nil {
		return err
	}

	w, err := c.Writer(ctx, typ)

	if err != nil {
		return err
	}

	_, err = io.Copy(w, r)

	if err != nil {
		return fmt.Errorf("failed to io.Copy: %w", err)
	}

	err = w.Close()

	return err
}

func main() {
	log.SetFlags(0)
	err := run()

	if err != nil {
		log.Fatal(err)
	}
}

// run starts a http.Server for the passed in address
// with all requests handled by echoServer.
func run() error {
	if len(os.Args) < 2 {
		return errors.New("please provide an address to listen on as the first argument")
	}

	l, err := net.Listen("tcp", os.Args[1])

	if err != nil {
		return err
	}

	log.Printf("listening on ws://%v", l.Addr())

	s := &http.Server{
		Handler:      echoServer{logf: log.Printf},
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	errc := make(chan error, 1)

	go func() {
		errc <- s.Serve(l)
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)

	select {
	case err := <-errc:
		log.Printf("failed to serve: %v", err)
	case sig := <-sigs:
		log.Printf("terminating: %v", sig)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	return s.Shutdown(ctx)
}
