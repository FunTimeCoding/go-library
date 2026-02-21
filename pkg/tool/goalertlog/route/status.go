package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlog/poller"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlog/store"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func Status(
	s *store.Store,
	p *poller.Poller,
) http.HandlerFunc {
	return func(
		w http.ResponseWriter,
		_ *http.Request,
	) {
		result := StatusResponse{
			TotalRecords: s.Count(),
		}
		lastPoll := p.LastPoll()

		if !lastPoll.IsZero() {
			result.LastPoll = lastPoll.Format(DateFormat)
		}

		w.Header().Set(constant.ContentType, constant.Object)
		errors.PanicOnError(json.NewEncoder(w).Encode(result))
	}
}
