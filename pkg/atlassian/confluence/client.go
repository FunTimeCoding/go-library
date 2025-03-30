package confluence

import (
	"context"
	treminio "github.com/ctreminiom/go-atlassian/v2/confluence"
	treminioV2 "github.com/ctreminiom/go-atlassian/v2/confluence/v2"
	kaos "github.com/essentialkaos/go-confluence/v6"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic_client"
	virtomize "github.com/virtomize/confluence-go-api"
)

type Client struct {
	host   string
	labels []string

	context context.Context

	basic *basic_client.Client

	kaos       *kaos.API
	virtomize  *virtomize.API
	treminio   *treminio.Client
	treminioV2 *treminioV2.Client
}
