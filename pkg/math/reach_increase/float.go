package reach_increase

func Float(
	past float64,
	now float64,
	threshold float64,
) bool {
	if past < threshold && now >= threshold {
		return true
	}

	return false
}
