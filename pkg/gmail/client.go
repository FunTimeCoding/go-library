package gmail

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/oauth/callback"
	"golang.org/x/oauth2"
	"google.golang.org/api/gmail/v1"
)

type Client struct {
	context    context.Context
	directory  string
	callback   *callback.Server
	credential *oauth2.Config
	service    *gmail.Service
}
