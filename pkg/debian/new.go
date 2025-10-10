package debian

import (
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"github.com/funtimecoding/go-library/pkg/system/join"
)

func New() *Client {
	h := system.Home()
	return &Client{
		home:          h,
		workDirectory: join.Absolute(h, constant.DownloadsPath),
	}
}
