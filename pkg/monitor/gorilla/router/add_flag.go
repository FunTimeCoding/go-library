package router

import "github.com/funtimecoding/go-library/pkg/monitor/gorilla/router/flag"

func (r *Router) AddFlag(f *flag.Flag) {
	r.Flag = append(r.Flag, f)
}
