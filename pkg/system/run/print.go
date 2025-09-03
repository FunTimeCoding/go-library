package run

import (
	"fmt"
	"strings"
)

func (r *Run) Print() {
	if r.Error != nil {
		fmt.Printf("Error: %s\n", r.Error)
	}

	if s := strings.TrimSpace(r.OutputString); s != "" {
		fmt.Printf("Stdout:\n%s\n", s)
	}

	if s := strings.TrimSpace(r.ErrorString); s != "" {
		fmt.Printf("Stderr:\n%s\n", s)
	}
}
