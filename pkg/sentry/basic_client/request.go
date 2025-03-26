package basic_client

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/strings/join/key_value"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (c *Client) Request(path string) string {
	client := web.Client(true)
	request, e := http.NewRequest(
		web.GetMethod,
		fmt.Sprintf("https://%s/api/0%s", c.host, path),
		nil,
	)
	errors.PanicOnError(e)
	request.Header.Add(web.ContentTypeHeader, web.ObjectContentType)
	request.Header.Add(
		web.AuthorizationHeader,
		key_value.Space(web.BearerPrefix, c.token),
	)
	request.Header.Add(web.AcceptHeader, web.ObjectContentType)
	response, doFail := client.Do(request)
	errors.PanicOnError(doFail)
	fmt.Println(request)
	fmt.Println(response)

	return web.ReadString(response)
}
