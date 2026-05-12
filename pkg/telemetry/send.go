package telemetry

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c *Client) send(
	tool string,
	surface string,
	actor string,
	outcome string,
	durationMillisecond int64,
) {
	body := map[string]any{
		"tool":    tool,
		"surface": surface,
		"actor":   actor,
		"outcome": outcome,
	}

	if durationMillisecond > 0 {
		body["duration_ms"] = durationMillisecond
	}

	encoded, e := json.Marshal(body)

	if e != nil {
		return
	}

	response, e := c.httpClient.Post(
		fmt.Sprintf("%s/api/events", c.base),
		"application/json",
		bytes.NewReader(encoded),
	)

	if e != nil {
		return
	}

	if f := response.Body.Close(); f != nil {
		return
	}
}
