package writer_mock

import "net/http"

type Mock struct {
	Headers http.Header
}
