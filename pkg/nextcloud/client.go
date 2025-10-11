package nextcloud

import (
	"github.com/funtimecoding/go-library/pkg/nextcloud/basic"
	"github.com/funtimecoding/go-library/pkg/web/authoring"
)

type Client struct {
	basic     *basic.Client
	authoring *authoring.Client
}
