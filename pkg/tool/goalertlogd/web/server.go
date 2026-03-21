package web

import (
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/poller"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/store"
)

type Server struct {
	store  *store.Store
	poller *poller.Poller
}
