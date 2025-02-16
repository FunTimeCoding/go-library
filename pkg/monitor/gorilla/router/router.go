package router

import (
	"github.com/funtimecoding/go-library/pkg/monitor/gorilla/router/client"
	"github.com/funtimecoding/go-library/pkg/monitor/gorilla/router/flag"
)

type Router struct {
	Client []*client.Client
	Flag   []*flag.Flag
}
