package notation_tag

import (
	"github.com/funtimecoding/go-library/pkg/strings/split"
	"github.com/funtimecoding/go-library/pkg/structs/constant"
	"reflect"
)

func Read(r reflect.StructField) []string {
	return split.Comma(r.Tag.Get(constant.NotationKey))
}
