package habitica

import (
	"bytes"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/system"
	"io"
	"net/http"
)

func (c *Client) do(
	method string,
	path string,
	body any,
) *http.Response {
	var reader io.Reader

	if body != nil {
		reader = bytes.NewReader(notation.Marshal(body))
	}

	r, e := http.NewRequest(method, join.Empty(c.baseURL, path), reader)
	errors.PanicOnError(e)
	r.Header.Set(userHeader, c.userID)
	r.Header.Set(tokenHeader, c.token)
	r.Header.Set("Content-Type", "application/json")
	result, f := c.http.Do(r)
	errors.PanicOnError(f)

	if result.StatusCode >= http.StatusBadRequest {
		b := system.ReadAll(result.Body)
		panic(
			fmt.Sprintf(
				"habitica %s %s: %d %s",
				method,
				path,
				result.StatusCode,
				string(b),
			),
		)
	}

	return result
}
