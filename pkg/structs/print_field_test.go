package structs

import (
	"reflect"
	"testing"
)

func TestPrintField(t *testing.T) {
	type Test struct {
		Name string `json:"name"`
	}
	PrintField(reflect.TypeOf(Test{}).Field(0))
}
