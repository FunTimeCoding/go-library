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

	if false {
		fmt.Printf("Chats: %s\n", c.Chats())
	}

	if false {
		fmt.Printf("Memories: %s\n", c.Memories())
	}

	if false {
		fmt.Printf("Notes: %s\n", c.Notes())
	}

	if true {
		fmt.Printf("Prompts: %s\n", c.Prompts())
	}
}
