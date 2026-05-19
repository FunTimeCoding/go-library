package service

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/client"
)

type Service struct {
	Store  *store.Store
	memory client.Client
}
