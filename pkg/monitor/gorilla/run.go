package gorilla

import (
	"flag"
	"log"
	"net/http"
)

func Run() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(*addressFlag, nil))
}
