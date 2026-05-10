package habitica

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/habitica/constant"
	"github.com/funtimecoding/go-library/pkg/habitica/response"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	web "github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/funtimecoding/go-library/pkg/web/detail_error"
	"io"
	"net/http"
)

func (c *Client) do(
	method string,
	path string,
	body any,
) (*http.Response, error) {
	var reader io.Reader

	if body != nil {
		reader = bytes.NewReader(notation.Marshal(body))
	}

	r, e := http.NewRequest(method, join.Empty(c.base, path), reader)

	if e != nil {
		return nil, e
	}

	r.Header.Set(constant.UserHeader, c.userIdentifier)
	r.Header.Set(constant.TokenHeader, c.token)
	r.Header.Set(web.ContentType, web.Object)
	result, f := c.client.Do(r)

	if f != nil {
		return nil, f
	}

	if result.StatusCode >= http.StatusBadRequest {
		b, g := io.ReadAll(result.Body)

		if h := result.Body.Close(); h != nil {
			return nil, h
		}

		if g != nil {
			return nil, fmt.Errorf(
				"habitica %s %s: %d (body unreadable: %w)",
				method,
				path,
				result.StatusCode,
				g,
			)
		}

		var r response.Error

		if json.Unmarshal(b, &r) == nil && r.Message != "" {
			return nil, detail_error.New(
				r.Message,
				result.Status,
			)
		}

		return nil, fmt.Errorf("%s", result.Status)
	}

	return result, nil
}
