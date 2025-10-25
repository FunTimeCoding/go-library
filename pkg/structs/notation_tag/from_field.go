package notation_tag

import "reflect"

func FromField(f reflect.StructField) *Tag {
	v := Read(f)

	if len(v) == 0 {
		return nil
	}

	return New(v)
}
