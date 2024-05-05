package go_mod

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/system"
)

func DowngradeDependencies(v []string) {
	for _, element := range v {
		fmt.Printf("Downgrade: %s\n", element)
		system.Run(constant.Go, constant.Get, element)
	}

	if len(v) > 0 {
		Tidy()
	}
}
