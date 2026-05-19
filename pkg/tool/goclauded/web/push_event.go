package web

import (
	"bytes"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"maragu.dev/gomponents"
	"net/http"
)

func pushEvent(
	w http.ResponseWriter,
	name string,
	node gomponents.Node,
) {
	var b bytes.Buffer
	errors.PanicOnError(node.Render(&b))
	_, e := fmt.Fprintf(w, "event: %s\ndata: %s\n\n", name, b.String())
	errors.PanicOnError(e)
}
