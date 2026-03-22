package example

import "example/inner"

func FlaggedNew() *inner.MyStruct {
	return new(inner.MyStruct) // want `use a constructor function instead of new\(inner\.MyStruct\)`
}
