package route

import (
	"github.com/funtimecoding/go-library/pkg/tool/goalertlog/poller"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlog/store"
)

type Handler struct {
	store  *store.Store
	poller *poller.Poller
}
