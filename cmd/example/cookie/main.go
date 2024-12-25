package main

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/network"
	"github.com/funtimecoding/go-library/pkg/text/multi_line"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/funtimecoding/go-library/pkg/web/request_context"
	"net/http"
	"sync"
)

const SessionCookie = "session"

var sessionStore = struct {
	sync.Mutex
	sessions map[string]string
}{
	sessions: make(map[string]string),
}

func main() {
	m := http.NewServeMux()
	m.HandleFunc(
		"/",
		func(
			w http.ResponseWriter,
			e *http.Request,
		) {
			c := request_context.New(w, e)
			l := multi_line.New()

			if s := c.Cookie(SessionCookie); s != nil {
				l.Format("session: %s", s.Value)
			} else {
				l.Add("no session")
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
			address := network.SplitHost(e.RemoteAddr)
			l := multi_line.New()
			l.Format("client: %s", address)
			s := c.Cookie(SessionCookie)

			if address == network.LocalhostAddressString {
				if s == nil {
					identifier := web.GenerateSession()
					sessionStore.Lock()
					sessionStore.sessions[identifier] = address
					sessionStore.Unlock()

					s = c.SetCookie(SessionCookie, identifier)
					l.Add("cookie created")
				} else {
					l.Add("cookie exists")
				}
			}

			if s != nil {
				l.Format("session: %s", s.Value)
			} else {
				l.Add("no session")
			}

			web.WriteOkay(w, l.Render())
		},
	)
	m.HandleFunc(
		"/logout",
		func(
			w http.ResponseWriter,
			e *http.Request,
		) {
			c := request_context.New(w, e)
			l := multi_line.New()

			if s := c.Cookie(SessionCookie); s != nil {
				sessionStore.Lock()
				delete(sessionStore.sessions, s.Value)
				sessionStore.Unlock()
				c.UnsetCookie(SessionCookie)
				l.Add("cookie removed")
			} else {
				l.Add("no session")
			}

			web.WriteOkay(w, l.Render())
		},
	)
	errors.PanicOnError(http.ListenAndServe(":8080", m))
}
