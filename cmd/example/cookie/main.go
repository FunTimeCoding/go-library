package main

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/text/multi_line"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/funtimecoding/go-library/pkg/web/authenticator"
	"github.com/funtimecoding/go-library/pkg/web/location"
	"github.com/funtimecoding/go-library/pkg/web/request_context"
	"net/http"
)

func main() {
	a := authenticator.New()
	m := http.NewServeMux()
	m.HandleFunc(
		location.Root,
		func(
			w http.ResponseWriter,
			e *http.Request,
		) {
			c := request_context.New(w, e)
			l := multi_line.New()

			if a.LoggedIn(c) {
				l.Add("logged in")
			} else {
				l.Add("not logged in")
			}

			c.WriteOkay(l.Render())
		},
	)
	m.HandleFunc(
		location.Login,
		func(
			w http.ResponseWriter,
			e *http.Request,
		) {
			c := request_context.New(w, e)
			a.AddressLogin(c)
			c.Redirect(location.Root)
		},
	)
	m.HandleFunc(
		location.Logout,
		func(
			w http.ResponseWriter,
			e *http.Request,
		) {
			c := request_context.New(w, e)
			a.Logout(c)
			c.Redirect(location.Root)
		},
	)
	errors.PanicOnError(http.ListenAndServe(web.ListenAddress, m))
}
