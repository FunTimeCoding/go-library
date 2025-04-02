package option

func Compact() *Whitespace {
	result := New()
	result.NewlineAtEnd = false
	result.AllowedBlankLines = 0

	return result
}
