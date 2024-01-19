package notation

func Format(object any) string {
	return Encode(object, true)
}
