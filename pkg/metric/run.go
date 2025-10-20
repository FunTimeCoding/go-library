package metric

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/web"
	"sync"
)

func (s *Server) Run(g *sync.WaitGroup) {
	if g == nil {
		g = NewWaitGroup()
	}

	if s.verbose {
		fmt.Printf("start metric server on %d\n", s.port)
	}

	g.Add(1)
	defer g.Done()
	web.ServeAsynchronous(s.server)
}
