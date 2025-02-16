package flag

import "github.com/funtimecoding/go-library/pkg/monitor/gorilla/router/client"

type Flag struct {
	Identifier string
	By         []*client.Client
}
