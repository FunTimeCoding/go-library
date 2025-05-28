package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/open_webui"
)

func main() {
	c := open_webui.NewEnvironment()

	if false {
		fmt.Printf("Models: %s\n", c.Models())
	}

	if true {
		fmt.Printf("Chats: %s\n", c.Chats())
	}
}
