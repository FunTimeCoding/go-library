package example

import "fmt"

func WrongError() {
	x := fmt.Errorf("test") // want `variable x of type error should be named e`
	_ = x
}
