package confluence

import (
	kaos "github.com/essentialkaos/go-confluence/v6"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic_client"
	virtomize "github.com/virtomize/confluence-go-api"
)

type Client struct {
	host      string
	virtomize *virtomize.API
	kaos      *kaos.API
	basic     *basic_client.Client
	labels    []string
}
