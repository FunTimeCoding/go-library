package writer_mock

import "net/http"

func New() *Mock {
	return &Mock{
		Headers: make(http.Header),
	}
}
