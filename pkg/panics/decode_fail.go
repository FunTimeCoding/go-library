package panics

import "fmt"

func DecodeFail(k string) {
	panic(fmt.Errorf("decode fail: %s", k))
}
