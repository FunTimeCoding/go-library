package gorilla

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/monitor/gorilla/router"
	"github.com/funtimecoding/go-library/pkg/web/location"
	"log"
	"net/http"
)

func Run(address string) {
	r := router.New()
	http.HandleFunc(location.Monitor, r.Monitor)
	http.HandleFunc(location.Echo, echo)
	http.HandleFunc(location.Root, home)
	log.Printf("listen on %s\n", address)
	errors.PanicOnError(http.ListenAndServe(address, nil))
}
