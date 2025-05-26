package basic_client

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (c *Client) Propfind() {
	req := web.NewPropfind(c.fileRoot)

	if false {
		// WebDAV is XML
		req.Header.Set(web.AcceptHeader, web.ObjectContentType)
	}

	req.SetBasicAuth(c.user, c.password)
	res := web.Send(web.Client(true), req)
	defer errors.LogClose(res.Body)

	switch res.StatusCode {
	case http.StatusMultiStatus:
		fmt.Println("success")

		if false {
			// A lot of XML
			body := web.ReadString(res)
			fmt.Println("response body:", body)
		}
	case http.StatusUnauthorized:
		fmt.Println("unauthorized")
	default:
		fmt.Printf("unexpected status: %d\n", res.StatusCode)
	}
}
