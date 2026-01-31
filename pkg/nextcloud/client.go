package nextcloud

import (
	"github.com/funtimecoding/go-library/pkg/nextcloud/basic"
	"github.com/funtimecoding/go-library/pkg/web/author"
)

type Client struct {
	basic  *basic.Client
	author *author.Client
}
