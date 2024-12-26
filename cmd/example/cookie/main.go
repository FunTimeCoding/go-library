package main

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/text/multi_line"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/funtimecoding/go-library/pkg/web/authenticator"
	"github.com/funtimecoding/go-library/pkg/web/request_context"
	"net/http"
)

func main() {
	a := authenticator.New()
	m := http.NewServeMux()
	m.HandleFunc(
		"/",
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

			web.WriteOkay(w, l.Render())
		},
	)
	m.HandleFunc(
		"/login",
		func(
			w http.ResponseWriter,
			e *http.Request,
		) {
			c := request_context.New(w, e)
			a.AddressLogin(c)
			c.Redirect("/")
		},
	)
	m.HandleFunc(
		"/logout",
		func(
			w http.ResponseWriter,
			e *http.Request,
		) {
			c := request_context.New(w, e)
			a.Logout(c)
			c.Redirect("/")
		},
	)
	errors.PanicOnError(http.ListenAndServe(":8080", m))
}
