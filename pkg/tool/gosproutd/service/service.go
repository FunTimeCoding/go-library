package service

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/tool/gosproutd/store"
)

type Service struct {
	store    *store.Store
	notifier face.EventNotifier
}
