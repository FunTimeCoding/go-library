package chroma

import (
	"context"
	"github.com/amikos-tech/chroma-go/pkg/api/v2"
	"github.com/funtimecoding/go-library/pkg/errors"
)

// TODO: For gomonitor, allow configuring to monitor files like ~/Library/Logs/cron.log and show if its newer than X minutes (good) and if its less, raise as a warning/error

func New() *Client {
	client, e := v2.NewHTTPClient()
	errors.PanicOnError(e)

	return &Client{context: context.Background(), client: client}
}
