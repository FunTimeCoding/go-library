package floats

func Reverse(v []float64) {
	last := len(v) - 1

	for i := 0; i < len(v)/2; i++ {
		v[i], v[last-i] = v[last-i], v[i]
	}
}
