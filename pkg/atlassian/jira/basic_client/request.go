package basic_client

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (c *Client) Request(path string) (int, string) {
	client := web.Client(true)
	request, e := http.NewRequest(
		web.GetMethod,
		fmt.Sprintf("%s%s", c.locator, path),
		nil,
	)
	errors.PanicOnError(e)
	request.SetBasicAuth(c.user, c.password)
	response, doFail := client.Do(request)
	errors.PanicOnError(doFail)

	return response.StatusCode, web.ReadString(response)
}
