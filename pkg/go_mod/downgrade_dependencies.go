package go_mod

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/system"
)

func DowngradeDependencies(v []string) {
	fmt.Printf("Downgrades: %+v\n", v)

	for _, element := range v {
		fmt.Printf("Downgrade: %s\n", element)
		system.Run("go", "get", element)
	}

	if len(v) > 0 {
		Tidy()
	}
}
