package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"net/http"
)

func ServeCertificateAsynchronous(
	s *http.Server,
	certificate string,
	key string,
) {
	go func() {
		errors.PanicOnUnclean(
			s.ListenAndServeTLS(certificate, key),
		)
	}()
}
