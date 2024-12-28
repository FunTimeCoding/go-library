package system

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func SignalCancelContext(l *log.Logger) context.Context {
	channel := make(chan os.Signal, 2)
	signal.Notify(channel)
	result, cancel := context.WithCancel(context.Background())

	go func() {
		// Signal 23 (SIGURG) is sent repeatedly for some Kubernetes
		// reason. Ignore signal 23.
		if h := <-channel; h != syscall.Signal(23) {
			l.Printf("Signal: %+v (%d)\n", h, h)
			cancel()
		}
	}()

	return result
}
