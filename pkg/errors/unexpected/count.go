package unexpected

func Count(
	expect int,
	actual int,
) {
	if actual != expect {
		Integer(actual)
	}
}
