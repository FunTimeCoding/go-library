package basic

import (
	"bytes"
	"encoding/json"
	"fmt"
	web "github.com/funtimecoding/go-library/pkg/web/constant"
	"io"
	"net/http"
)

func (c *Client) Post(
	path string,
	body any,
) ([]byte, error) {
	b, e := json.Marshal(body)

	if e != nil {
		return nil, e
	}

	r, f := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/%s", c.base, path),
		bytes.NewReader(b),
	)

	if f != nil {
		return nil, f
	}

	r.Header.Set(web.Accept, web.Object)
	r.Header.Set(web.ContentType, web.Object)


	s, g := c.client.Do(r)

	if g != nil {
		return nil, g
	}

	result, h := io.ReadAll(s.Body)
	i := s.Body.Close()

	if h != nil {
		return nil, h
	}

	if i != nil {
		return nil, i
	}

	if s.StatusCode >= http.StatusBadRequest {
		return nil, parseDetail(result, s.Status)
	}

	return result, nil
}
