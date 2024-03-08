package markup

func Render(input any) string {
	return string(Clean(Encode(input)))
}
