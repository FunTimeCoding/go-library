package example

import "fmt"

func WrongErrorMulti() {
	err := fmt.Errorf("test") // want `variable err of type error should be named e`
	_ = err
}
