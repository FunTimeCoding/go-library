package example

import "fmt"

func ErrorChain() {
	e := fmt.Errorf("first")
	x := fmt.Errorf("second") // want `variable x of type error should be named f`
	_ = e
	_ = x
}
