package raid

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/tool/goraidd/client"
)

type Client struct {
	context context.Context
	client  *client.ClientWithResponses
}
