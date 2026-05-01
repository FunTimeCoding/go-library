package basic

import (
	"io"
	"net/http"
)

func (c *Client) DeleteV2(l string) (string, error) {
	r, e := http.NewRequest(http.MethodDelete, l, nil)

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

	return string(b), nil
}
