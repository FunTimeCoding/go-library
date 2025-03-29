package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/opsgenie/opsgenie-go-sdk-v2/user"
)

func User(c *client.Config) *user.Client {
	result, e := user.NewClient(c)
	errors.PanicOnError(e)

	return result
}
