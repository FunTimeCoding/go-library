package kestra

import (
	"github.com/funtimecoding/go-library/pkg/strings/join/key_value"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/kestra-io/client-sdk/go-sdk"
)

func New(
	host string,
	token string,
) *Client {
	// Currently locked behind EE: https://kestra.io/docs/enterprise/auth/api-tokens
	c := kestra_api_client.NewConfiguration()
	c.Host = host
	c.DefaultHeader[constant.AuthorizationHeader] = key_value.Space(
		constant.BearerPrefix,
		token,
	)

	return &Client{client: kestra_api_client.NewAPIClient(c)}
}
