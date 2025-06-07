package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/gmail"
	"github.com/funtimecoding/go-library/pkg/gmail/constant"
)

func main() {
	c := gmail.NewEnvironment().Load()
	r := c.Unread()
	fmt.Printf("Unread (%d):\n", len(r.Messages))

	for _, m := range r.Messages {
		var subject string

		for _, h := range c.Message(m).Payload.Headers {
			if h.Name == constant.SubjectHeader {
				subject = h.Value

				break
			}
		}

		fmt.Printf("- %s\n", subject)
	}
}
