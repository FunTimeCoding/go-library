package example

import "fmt"

func WrongErrorMulti() {
	fail := fmt.Errorf("test") // want `variable fail of type error should be named e`
	_ = fail
}
