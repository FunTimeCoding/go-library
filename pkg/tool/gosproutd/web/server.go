package web

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/tool/gosproutd/store"
	"github.com/funtimecoding/go-library/pkg/web/view"
)

type Server struct {
	store    *store.Store
	notifier face.EventNotifier
	view     *view.View
}
