package chat_request

import "fmt"

func (r *Request) Print() {
	for _, m := range r.request.Messages {
		fmt.Printf("%s: %s\n", m.Role, m.Content)
	}
}
