package writer_mock

import "net/http"

type Mock struct {
	Headers http.Header
}

func (m *Mock) Header() http.Header {
	return m.Headers
}

func (m *Mock) Write([]byte) (int, error) {
	panic("implement me")
}

func (m *Mock) WriteHeader(int) {
	panic("implement me")
}
