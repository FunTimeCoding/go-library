package habitica

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"net/http"
)

func (c *Client) decode(
	r *http.Response,
	result any,
) {
	var envelope struct {
		Data json.RawMessage `json:"data"`
	}
	e := json.NewDecoder(r.Body).Decode(&envelope)
	errors.PanicOnError(e)
	f := json.Unmarshal(envelope.Data, result)
	errors.PanicOnError(f)
}
