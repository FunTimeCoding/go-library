package debian

import (
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/constant"
)

func New() *Client {
	h := system.Home()
	return &Client{
		home:          h,
		workDirectory: system.Join(h, constant.DownloadsPath),
	}
}
