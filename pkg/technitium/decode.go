package technitium

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/technitium/envelope"
	"io"
	"net/http"
)

func (c *Client) decode(
	r *http.Response,
	result any,
) error {
	b, e := io.ReadAll(r.Body)

	if f := r.Body.Close(); f != nil {
		return f
	}

	if e != nil {
		return e
	}

	var v envelope.Envelope

	if f := json.Unmarshal(b, &v); f != nil {
		return f
	}

	if v.Status != "ok" {
		return fmt.Errorf("technitium: %s", v.Message)
	}

	return json.Unmarshal(v.Payload, result)
}
