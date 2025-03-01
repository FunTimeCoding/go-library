package router

import "github.com/funtimecoding/go-library/pkg/monitor/gorilla/router/flag"

func (r *Router) FlagByIdentifier(i string) *flag.Flag {
	for _, l := range r.Flag {
		if l.Identifier == i {
			return l
		}
	}

	return nil
}
