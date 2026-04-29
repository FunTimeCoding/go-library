package server

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/raid"
	"github.com/funtimecoding/go-library/pkg/time"
	generated "github.com/funtimecoding/go-library/pkg/tool/goraidd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) GetLogs(
	w http.ResponseWriter,
	_ *http.Request,
	params generated.GetLogsParams,
) {
	all := s.store.Fights()
	var fights []raid.Fight

	for _, f := range all {
		if params.Start != nil && f.Timestamp.Before(*params.Start) {
			continue
		}

		if params.End != nil && f.Timestamp.After(*params.End) {
			continue
		}

		fights = append(fights, f)
	}

	total := len(fights)
	offset := 0
	limit := 50

	if params.Offset != nil {
		offset = *params.Offset
	}

	if params.Limit != nil {
		limit = *params.Limit
	}

	if offset > total {
		offset = total
	}

	end := offset + limit

	if end > total {
		end = total
	}

	page := fights[offset:end]
	var result []generated.LogResponse

	for _, f := range page {
		result = append(
			result,
			generated.LogResponse{
				FileName:    f.Filename,
				Time:        f.Timestamp.Format(time.DateSecond),
				Duration:    fmt.Sprintf("%dms", f.DurationMS),
				MapId:       f.MapID,
				PlayerCount: f.AlliedCount,
				Players:     nil,
			},
		)
	}

	web.EncodeNotation(
		w,
		generated.LogsResponse{
			Total: total,
			Logs:  result,
		},
	)
}
