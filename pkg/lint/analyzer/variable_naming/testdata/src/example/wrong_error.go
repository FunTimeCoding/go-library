package example

import "fmt"

func WrongErrorSingleLetter() {
	x := fmt.Errorf("test") // want `variable x of type error should be named e`
	_ = x
}
