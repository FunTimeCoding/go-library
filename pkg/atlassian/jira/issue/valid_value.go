package issue

func ValidValue(s string) bool {
	if s == NilValue || s == UnknownValue || s == UnknownField {
		return false
	}

	return true
}
