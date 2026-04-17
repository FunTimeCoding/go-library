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
) {
	logs := log.NewSlice(
		gw2.ParseLogs(
			system.ReadBytes(h.logCachePath, gw2Constant.LogFile),
			false,
		),
	)
	var result []generated.LogResponse

	for _, l := range logs {
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
	errors.PanicOnError(json.NewEncoder(w).Encode(result))
}
