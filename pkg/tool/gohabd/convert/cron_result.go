package convert

import (
	"github.com/funtimecoding/go-library/pkg/habitica"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/generated/server"
)

func CronResult(r habitica.CronResult) server.CronResult {
	result := server.CronResult{
		RolledOver: r.RolledOver,
		LastCron:   r.LastCron,
	}

	if r.Before != nil {
		s := statsPointer(*r.Before)
		result.Before = s
	}

	if r.After != nil {
		s := statsPointer(*r.After)
		result.After = s
	}

	return result
}
