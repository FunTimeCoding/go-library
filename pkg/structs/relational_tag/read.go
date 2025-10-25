package relational_tag

import (
	"github.com/funtimecoding/go-library/pkg/strings/split"
	"github.com/funtimecoding/go-library/pkg/structs/constant"
	"reflect"
)

func Read(f reflect.StructField) []string {
	return split.Semicolon(f.Tag.Get(constant.RelationalKey))
}
