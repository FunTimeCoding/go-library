package markup

func Render(a any) string {
	return string(Clean(Encode(a)))
}
