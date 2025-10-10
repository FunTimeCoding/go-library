package basic_client

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/funtimecoding/go-library/pkg/web/locator"
)

func (c *Client) GetV2(path string) string {
	l := locator.NewHost(c.host).Base(constant.Base).Path(path).String()
	r := web.NewGet(l)
	r.SetBasicAuth(c.user, c.token)
	response := web.Send(web.Client(true), r)
	result := web.ReadString(response)

	if response.StatusCode >= 400 {
		fmt.Printf("Locator: %s\n", l)
		fmt.Printf("Status: %d\n", response.StatusCode)
		web.PrintHeader(response.Header)
		fmt.Printf("Body: %s\n", result)
	}

	return result
}
