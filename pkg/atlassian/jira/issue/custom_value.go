package issue

import "fmt"

func (i *Issue) CustomValue(fieldName string) string {
	field := i.fieldMap.ByName(fieldName)

	if field == nil {
		return UnknownField
	}

	for k, v := range i.Raw.Fields.Unknowns {
		if k == field.ID {
			if v == nil {
				return NilValue
			}

			return fmt.Sprintf("%+v", v)
		}
	}

	return UnknownValue
}
