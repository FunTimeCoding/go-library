package panics

import "fmt"

func Unexpected(k string) {
	panic(fmt.Errorf("unexpected: %s", k))
}
