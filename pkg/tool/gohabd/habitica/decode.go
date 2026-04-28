package habitica

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/system"
	"net/http"
)

func (c *Client) decode(
	r *http.Response,
	result any,
) {
	var envelope struct {
		Payload json.RawMessage `json:"data"`
	}
	notation.DecodeBytesStrict(
		system.ReadAll(r.Body),
		&envelope,
		false,
	)
	notation.DecodeBytesStrict(envelope.Payload, result, false)
}
