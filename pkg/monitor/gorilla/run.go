package gorilla

import (
	"github.com/funtimecoding/go-library/pkg/monitor/gorilla/router"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/funtimecoding/go-library/pkg/web/location"
	"log"
	"net/http"
)

func Run(address string) {
	r := router.New()
	m := http.NewServeMux()
	m.HandleFunc(location.Monitor, r.Monitor)
	m.HandleFunc(location.Echo, echo)
	m.HandleFunc(location.Root, home)
	log.Printf("listen on %s\n", address)
	web.ListenAddress(m, address)
}
