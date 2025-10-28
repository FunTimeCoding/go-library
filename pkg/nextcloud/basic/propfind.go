package basic

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func (c *Client) Propfind() {
	r := web.NewPropfind(c.fileRoot)

	if false {
		// WebDAV is XML
		r.Header.Set(constant.Accept, constant.Object)
	}

	r.SetBasicAuth(c.user, c.password)
	s := web.Send(web.Client(), r)
	defer errors.LogClose(s.Body)

	switch s.StatusCode {
	case http.StatusMultiStatus:
		fmt.Println("success")

		if false {
			// A lot of XML
			fmt.Println("response body:", web.ReadString(s))
		}
	case http.StatusUnauthorized:
		fmt.Println("unauthorized")
	default:
		fmt.Printf("unexpected status: %d\n", s.StatusCode)
	}
}
