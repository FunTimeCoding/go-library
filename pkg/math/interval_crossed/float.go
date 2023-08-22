package interval_crossed

func Float(
	interval float64,
	past float64,
	now float64,
) bool {
	return int(past/interval) != int(now/interval)
}
