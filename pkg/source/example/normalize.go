package example

func normalize(name string) string {
	if IsValid(name) {
		return name
	}

	return "default"
}
