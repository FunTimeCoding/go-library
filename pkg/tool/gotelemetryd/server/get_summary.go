package server

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gotelemetryd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) GetSummary(
	w http.ResponseWriter,
	_ *http.Request,
	params server.GetSummaryParams,
) {
	since := ""
	until := ""
	groupBy := "tool"

	if params.Since != nil {
		since = *params.Since
	}

	if params.Until != nil {
		until = *params.Until
	}

	if params.GroupBy != nil {
		groupBy = string(*params.GroupBy)
	}

	rows, queryError := s.store.Summary(since, until, groupBy)
	errors.PanicOnError(queryError)
	var entries []server.SummaryEntry

	for _, r := range rows {
		entry := server.SummaryEntry{
			Tool:  r.Tool,
			Count: int(r.Count),
		}

		if r.Surface != "" {
			entry.Surface = &r.Surface
		}

		if r.Kind != "" {
			entry.Kind = &r.Kind
		}

		entries = append(entries, entry)
	}

	if entries == nil {
		entries = []server.SummaryEntry{}
	}

	web.EncodeNotation(w, entries)
}
