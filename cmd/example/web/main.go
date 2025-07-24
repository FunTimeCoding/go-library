package main

import (
	"github.com/funtimecoding/go-library/pkg/text/multi_line"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/funtimecoding/go-library/pkg/web/location"
	"github.com/funtimecoding/go-library/pkg/web/request_context"
	"net/http"
)

func main() {
	m := http.NewServeMux()
	m.HandleFunc(
		location.Root,
		func(
			w http.ResponseWriter,
			e *http.Request,
		) {
			c := request_context.New(w, e)
			c.SetLastLocation()
			l := multi_line.New()
			l.Add("Hello friend.")
			c.WriteOkay(l.Render())
		},
	)
	web.Listen(m)
}
