package web

func (s *Server) drain(c chan struct{}) {
	for {
		select {
		case <-c:
		default:
			return
		}
	}
}
