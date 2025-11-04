package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"log"
	"net/http"
	"net/http/httputil"
)

func VerboseTrip(
	t http.RoundTripper,
	q *http.Request,
) (*http.Response, error) {
	reqDump, requestFail := httputil.DumpRequestOut(q, true)
	errors.PanicOnError(requestFail)
	log.Printf("REQUEST:\n%s\n", string(reqDump))
	s, e := t.RoundTrip(q)

	if s != nil {
		respDump, responseFail := httputil.DumpResponse(s, true)
		errors.PanicOnError(responseFail)
		log.Printf("RESPONSE:\n%s\n", string(respDump))
	}

	return s, e
}
