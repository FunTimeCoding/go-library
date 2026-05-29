package server

import (
	"golang.org/x/sys/unix"
	"os"
	"os/signal"
	"sync"
)

func (s *Server) Run() error {
	var waitGroup sync.WaitGroup
	listenErrors := make(chan error, 1)

	for _, p := range s.processes {
		p.Spawn(s.environment.Build(), &waitGroup)
	}

	go func() {
		if e := s.Listen(); e != nil {
			listenErrors <- e
		}
	}()
	signals := make(chan os.Signal, 10)
	signal.Notify(signals, unix.SIGTERM, unix.SIGINT, unix.SIGHUP)
	allDone := make(chan struct{}, 1)
	go func() {
		waitGroup.Wait()
		allDone <- struct{}{}
	}()

	select {
	case <-allDone:
		return s.stopAll()
	case <-signals:
		return s.stopAll()
	case e := <-listenErrors:
		stopError := s.stopAll()

		if stopError != nil {
			return stopError
		}

		return e
	}
}
