package example

import "example/inner"

func FlaggedUnary() *inner.MyStruct {
	return &inner.MyStruct{} // want `use a constructor function instead of &inner\.MyStruct\{\}`
}
