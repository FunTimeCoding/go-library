package scheduler

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/robfig/cron/v3"
	"sync"
)

type Scheduler struct {
	cron     *cron.Cron
	schedule string
	task     func()
	logger   *logger.Logger
	reporter face.Reporter
	entry    cron.EntryID
	mutex    sync.Mutex
	running  bool
}
