package callback

import (
	"context"
	"net/http"
)

type Server struct {
	verbose    bool
	port       int
	context    context.Context
	server     *http.Server
	callbackCh chan string
	errorCh    chan error
}
