package basic

import (
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetV2(l string) (string, error) {
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

	body := string(b)

	if result.StatusCode >= 400 {
		return "", fmt.Errorf(
			"confluence GET %s: %d %s",
			l,
			result.StatusCode,
			body,
		)
	}

	return body, nil
}
