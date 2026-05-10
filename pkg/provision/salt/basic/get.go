package basic

import (
	"fmt"
	web "github.com/funtimecoding/go-library/pkg/web/constant"
	"io"
	"net/http"
)

func (c *Client) Get(path string) ([]byte, error) {
	r, e := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/%s", c.base, path),
		nil,
	)

	if e != nil {
		return nil, e
	}

	r.Header.Set(web.Accept, web.Object)
	s, f := c.client.Do(r)

	if f != nil {
		return nil, f
	}

	b, g := io.ReadAll(s.Body)
	h := s.Body.Close()

	if g != nil {
		return nil, g
	}

	if h != nil {
		return nil, h
	}

	if s.StatusCode >= http.StatusBadRequest {
		return nil, parseDetail(b, s.Status)
	}

	return b, nil
}
