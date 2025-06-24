package collect

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/bubbletea/model/monitor/fetch"
)

func Check() {
	for _, s := range fetch.List() {
		fmt.Printf("Command: %s\n", s)

		if items := fetch.Run(s); len(items) > 0 {
			for _, i := range items {
				fmt.Printf("  %s: %s\n", i.Identifier, i.Detail)
			}
		} else {
			fmt.Println("  No items")
		}
	}
}
