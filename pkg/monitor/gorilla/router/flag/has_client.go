package flag

import "github.com/funtimecoding/go-library/pkg/monitor/gorilla/router/client"

func (f *Flag) HasClient(v *client.Client) bool {
	for _, c := range f.By {
		if c == v {
			return true
		}
	}

	return false
}
