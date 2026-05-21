package web

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/face"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/service"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/web/conversations"
	"github.com/funtimecoding/go-library/pkg/web/view"
)

type Server struct {
	service       *service.Service
	notifier      face.Notifier
	conversations *conversations.Server
	view          *view.View
}
