package basic_client

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (c *Client) Get(path string) string {
	client := web.Client(true)
	request, e := http.NewRequest(
		web.GetMethod,
		fmt.Sprintf(
			"https://%s/loki/api/v1%s",
			c.host,
			path,
		),
		nil,
	)
	errors.PanicOnError(e)
	request.SetBasicAuth(c.user, c.password)
	response, doFail := client.Do(request)
	errors.PanicOnError(doFail)

	if false {
		fmt.Println(request)
		fmt.Println(response)
	}

	return web.ReadString(response)
}
