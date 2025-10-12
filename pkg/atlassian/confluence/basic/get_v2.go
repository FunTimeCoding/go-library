package basic

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) GetV2(l string) string {
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
