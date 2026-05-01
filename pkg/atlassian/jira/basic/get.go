package basic

import (
	"io"
	"net/http"
)

func (c *Client) Get(l string) (int, string, error) {
	r, e := http.NewRequest(http.MethodGet, l, nil)

	if e != nil {
		return 0, "", e
	}

	r.SetBasicAuth(c.user, c.token)
	result, f := http.DefaultClient.Do(r)

	if f != nil {
		return 0, "", f
	}

	b, g := io.ReadAll(result.Body)

	if h := result.Body.Close(); h != nil {
		return 0, "", h
	}

	if g != nil {
		return 0, "", g
	}

	return result.StatusCode, string(b), nil
}
