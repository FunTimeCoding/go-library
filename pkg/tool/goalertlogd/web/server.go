package web

import (
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/store"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/worker"
)

type Server struct {
	store  *store.Store
	worker *worker.Worker
}
