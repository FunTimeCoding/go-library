package log

import (
	"github.com/funtimecoding/go-library/pkg/gw2/log_manager"
	"time"
)

type Log struct {
	Accounts []string
	Time     time.Time

	Raw *log_manager.Log
}
