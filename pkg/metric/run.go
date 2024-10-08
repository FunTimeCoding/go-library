package metric

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/web"
	"sync"
)

func (s *Server) Run(g *sync.WaitGroup) {
	if s.verbose {
		fmt.Printf("metric server running on port %d\n", s.port)
	}

	g.Add(1)
	defer g.Done()
	web.ServeAsynchronous(s.server)
}
