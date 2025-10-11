package writer_mock

import "net/http"

func (m *Mock) Header() http.Header {
	return m.Headers
}
