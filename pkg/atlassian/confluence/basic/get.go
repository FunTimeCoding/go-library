package basic

import (
	"io"
	"net/http"
)

func (c *Client) Get(l string) (string, error) {
	r, e := http.NewRequest(http.MethodGet, l, nil)

	if e != nil {
		return "", e
	}

	r.SetBasicAuth(c.user, c.token)
	result, f := http.DefaultClient.Do(r)

	if f != nil {
		return "", f
	}

	b, g := io.ReadAll(result.Body)

	if h := result.Body.Close(); h != nil {
		return "", h
	}

	if g != nil {
		return "", g
	}

	if result.StatusCode >= http.StatusBadRequest {
		return "", parseDetail(b, result.Status)
	}

	return string(b), nil
}
