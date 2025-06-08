package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/brave"
	"github.com/funtimecoding/go-library/pkg/brave/constant"
)

func main() {
	for _, p := range brave.Profiles() {
		fmt.Printf("Profile: %s\n", p)

		if p == constant.Profile2 {
			brave.OpenProfile(p)
		}
	}
}
