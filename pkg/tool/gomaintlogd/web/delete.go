package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/constant"
	web "github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func (s *Server) delete(
	w http.ResponseWriter,
	r *http.Request,
) {
	errors.PanicOnError(
		s.store.Delete(
			strings.ToUnsignedIntegerStrict(
				r.URL.Query().Get(constant.Identifier),
			),
		),
	)
	w.Header().Set(web.ContentType, web.MarkupUnicode)
	w.WriteHeader(http.StatusOK)
}
