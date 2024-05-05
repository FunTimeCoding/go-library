package unexpected

func Count(
	expected int,
	actual int,
) {
	if actual != expected {
		Integer(actual)
	}
}
