package service

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/tool/gosproutd/store"
	"sync"
)

type Service struct {
	mutex    sync.Mutex
	store    *store.Store
	notifier face.EventNotifier
}
