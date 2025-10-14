package verbose_transport

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"log"
	"net/http"
	"net/http/httputil"
)

func (d *Transport) RoundTrip(r *http.Request) (*http.Response, error) {
	reqDump, requestFail := httputil.DumpRequestOut(r, true)
	errors.PanicOnError(requestFail)
	log.Printf("REQUEST:\n%s\n", string(reqDump))
	s, e := d.transport.RoundTrip(r)

	if s != nil {
		respDump, responseFail := httputil.DumpResponse(s, true)
		errors.PanicOnError(responseFail)
		log.Printf("RESPONSE:\n%s\n", string(respDump))
	}

	return s, e
}
