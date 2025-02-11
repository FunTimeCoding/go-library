package gorilla

import (
	"flag"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"
)

func RunClient() {
	flag.Parse()
	log.SetFlags(0)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	u := url.URL{Scheme: "ws", Host: *addressFlag, Path: "/echo"}
	log.Printf("connecting to %s", u.String())
	c, _, dialFail := websocket.DefaultDialer.Dial(
		u.String(),
		nil,
	)

	if dialFail != nil {
		log.Fatal("dial:", dialFail)
	}

	defer errors.LogClose(c)

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			if e := c.WriteMessage(
				websocket.TextMessage,
				[]byte(t.String()),
			); e != nil {
				log.Println("write:", e)

				return
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			if e := c.WriteMessage(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(
					websocket.CloseNormalClosure,
					"",
				),
			); e != nil {
				log.Println("write close:", e)

				return
			}

			select {
			case <-done:
			case <-time.After(time.Second):
			}

			return
		}
	}
}
