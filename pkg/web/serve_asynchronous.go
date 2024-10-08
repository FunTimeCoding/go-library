package web

import (
	"errors"
	"log"
	"net/http"
)

func ServeAsynchronous(s *http.Server) {
	go func() {
		if e := s.ListenAndServe(); e != nil {
			if !errors.Is(e, http.ErrServerClosed) {
				log.Panicf("listen: %s\n", e)
			}
		}
	}()
}
