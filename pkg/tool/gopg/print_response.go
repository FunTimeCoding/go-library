package gopg

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/notation"
	"io"
	"net/http"
)

func printResponse(r *http.Response) {
	defer errors.PanicClose(r.Body)
	b, e := io.ReadAll(r.Body)
	errors.PanicOnError(e)

	if r.StatusCode >= 400 {
		errors.Printf("error %d: %s", r.StatusCode, string(b))

		return
	}

	var v any
	errors.PanicOnError(json.Unmarshal(b, &v))
	fmt.Println(notation.MarshalIndent(v))
}
