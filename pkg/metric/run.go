package metric

import (
	"github.com/funtimecoding/go-library/pkg/web"
	"sync"
)

func (s *Server) Run(g *sync.WaitGroup) {
	g.Add(1)
	defer g.Done()
	web.ServeAsynchronous(s.server)
}
