package ssh

import (
	"bytes"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/separator"
	"strings"
)

func (c *Client) Run(command string) string {
	session, e := c.client.NewSession()
	errors.PanicOnError(e)
	defer func() {
		if f := session.Close(); f != nil && f.Error() != "EOF" {
			fmt.Println("Session error")
			errors.LogOnError(f)
		}
	}()
	var buffer bytes.Buffer
	session.Stdout = &buffer
	errors.PanicOnError(session.Run(command))

	return strings.TrimRight(buffer.String(), separator.Unix)
}
