package main

import (
	"github.com/funtimecoding/go-library/pkg/metric"
	"github.com/funtimecoding/go-library/pkg/system"
	"log"
)

func main() {
	m := metric.New(0, true)
	log.Println("starting")
	m.Run(nil)
	log.Println("started")
	system.KillSignalBlock()
	log.Println("stopping")
	m.Stop()
	log.Println("stopped")
}
