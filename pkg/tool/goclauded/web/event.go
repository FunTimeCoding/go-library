package web

import (
	"net/http"
	"time"
)

func (s *Server) event(
	w http.ResponseWriter,
	r *http.Request,
) {
	flusher, okay := w.(http.Flusher)

	if !okay {
		http.Error(
			w,
			"streaming not supported",
			http.StatusInternalServerError,
		)

		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	subscription := s.notifier.Subscribe()
	defer s.notifier.Unsubscribe(subscription)
	s.pushSections(w, flusher)

	for {
		select {
		case <-r.Context().Done():
			return
		case <-subscription:
			time.Sleep(1 * time.Second)
			s.drain(subscription)
			s.pushSections(w, flusher)
		}
	}
}
