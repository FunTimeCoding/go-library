package watcher

import (
	"github.com/funtimecoding/go-library/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/tool/gosproutd/store"
)

type Watcher struct {
	store         *store.Store
	notifier      face.EventNotifier
	recovery      *recovery.Recovery
	seedDirectory string
	stop          chan struct{}
}
