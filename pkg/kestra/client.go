package kestra

import (
	"context"
	"github.com/kestra-io/client-sdk/go-sdk"
)

type Client struct {
	context  context.Context
	client   *kestra_api_client.APIClient
	token    string
	user     string
	password string
}
