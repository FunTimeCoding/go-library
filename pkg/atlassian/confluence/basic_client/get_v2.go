package basic_client

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) GetV2(path string) string {
	locator := fmt.Sprintf(
		"%s://%s%s%s",
		c.scheme,
		c.host,
		constant.PathPrefix,
		path,
	)
	r := web.NewGet(locator)
	r.SetBasicAuth(c.user, c.token)
	response := web.Send(web.Client(true), r)
	result := web.ReadString(response)

	if response.StatusCode >= 400 {
		fmt.Printf("Location: %s\n", locator)
		fmt.Printf("Status: %d\n", response.StatusCode)
		web.PrintHeader(response.Header)
		fmt.Printf("Body: %s\n", result)
	}

	return result
}
