package habitica

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/habitica/response"
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

	var envelope response.Envelope

	if f := json.Unmarshal(b, &envelope); f != nil {
		return f
	}

	return json.Unmarshal(envelope.Payload, result)
}
