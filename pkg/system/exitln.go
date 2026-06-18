package system

import (
	"fmt"
	"os"
)

func Exitln(a ...any) {
	fmt.Println(a...)
	os.Exit(1)
}
