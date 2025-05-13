package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/scanner"
	"github.com/funtimecoding/go-library/pkg/generative/token"
	"github.com/funtimecoding/go-library/pkg/strings/join"
)

func Token() {
	s := scanner.New()
	fmt.Println("Paste text and press Ctrl+D to finish:")
	lines := s.Scan()

	if false {
		for _, l := range lines {
			fmt.Printf("Line: %+v\n", l)
		}
	}

	fmt.Printf("Token count: %d\n", token.Count(join.NewLine(lines)))
}
