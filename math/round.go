package math

import "math"

func Round(
	number float64,
	decimals int,
) float64 {
	var result float64

	switch decimals {
	case 1:
		result = math.Round(number*10) / 10
	case 2:
		result = math.Round(number*100) / 100
	case 3:
		result = math.Round(number*1000) / 1000
	case 4:
		result = math.Round(number*10000) / 10000
	case 5:
		result = math.Round(number*100000) / 100000
	case 6:
		result = math.Round(number*1000000) / 1000000
	case 7:
		result = math.Round(number*10000000) / 10000000
	case 8:
		result = math.Round(number*100000000) / 100000000
	default:
		result = math.Round(number)
	}

	return result
}
