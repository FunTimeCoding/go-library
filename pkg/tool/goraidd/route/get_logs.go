package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gw2"
	gw2Constant "github.com/funtimecoding/go-library/pkg/gw2/constant"
	"github.com/funtimecoding/go-library/pkg/gw2/log_manager/log"
	"github.com/funtimecoding/go-library/pkg/system"
	generated "github.com/funtimecoding/go-library/pkg/tool/goraidd/server"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func (h *Router) GetLogs(
	w http.ResponseWriter,
	_ *http.Request,
	params generated.GetLogsParams,
) {
	all := log.NewSlice(
		gw2.ParseLogs(
			system.ReadBytes(h.logCachePath, gw2Constant.LogFile),
			false,
		),
	)
	var logs []*log.Log

	for _, l := range all {
		if params.Start != nil && l.Time.Before(*params.Start) {
			continue
		}

		if params.End != nil && l.Time.After(*params.End) {
			continue
		}

		logs = append(logs, l)
	}

	total := len(logs)
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

	page := logs[offset:end]
	var result []generated.LogResponse

	for _, l := range page {
		result = append(
			result,
			generated.LogResponse{
				FileName:    l.Raw.FileName,
				Time:        l.Time.Format("2006-01-02 15:04:05"),
				Duration:    l.Raw.EncounterDuration,
				MapId:       l.Raw.MapID,
				PlayerCount: len(l.Raw.Players),
				Players:     l.Accounts,
			})
	}

	w.Header().Set(constant.ContentType, constant.Object)
	errors.PanicOnError(json.NewEncoder(w).Encode(generated.LogsResponse{
		Total: total,
		Logs:  result,
	}))
}
