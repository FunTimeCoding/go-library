package web

import (
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/store"
	"github.com/funtimecoding/go-library/pkg/web/palette"
	"github.com/funtimecoding/go-library/pkg/web/view"
)

type Server struct {
	store    *store.Store
	view     *view.View
	registry *palette.Registry
}
