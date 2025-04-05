package spy_writer

import "net/http"

func (s *Writer) Header() http.Header {
	return http.Header{}
}
