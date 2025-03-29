package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/opsgenie/opsgenie-go-sdk-v2/logs"
)

func Log(c *client.Config) *logs.Client {
	result, e := logs.NewClient(c)
	errors.PanicOnError(e)

	return result
}
