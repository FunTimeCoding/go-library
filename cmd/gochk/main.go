package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"runtime"
)

func main() {
	switch runtime.GOOS {
	case constant.Linux:
		fmt.Println("Linux")
	case constant.Darwin:
		fmt.Println("Darwin")
	}
}
