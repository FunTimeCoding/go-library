package model_context

func (s *Server) pushSummary(
	name string,
	body string,
) {
	if s.indexer == nil {
		return
	}

	if e := s.indexer.Push(name, body); e != nil {
		s.reporter.CaptureException(e)
	}
}
