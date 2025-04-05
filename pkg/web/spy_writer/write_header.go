package spy_writer

func (s *Writer) WriteHeader(statusCode int) {
	s.StatusCode = statusCode
}
