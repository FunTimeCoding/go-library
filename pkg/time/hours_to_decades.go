package time

func HoursToDecades(hours float64) float64 {
	return ((hours / 24) * 7 * 4 * 12) / (365 * 10)
}
