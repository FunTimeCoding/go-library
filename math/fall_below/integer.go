package fall_below

func Integer(
	past int,
	now int,
	threshold int,
) bool {
	if past >= threshold && now < threshold {
		return true
	}

	return false
}
