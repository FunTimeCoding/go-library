package nextcloud

import (
	"github.com/funtimecoding/go-library/pkg/nextcloud/basic_client"
	"github.com/funtimecoding/go-library/pkg/web/authoring"
)

type Client struct {
	basic     *basic_client.Client
	authoring *authoring.Client
}
