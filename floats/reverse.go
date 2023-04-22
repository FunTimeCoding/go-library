package floats

func Reverse(elements []float64) {
	last := len(elements) - 1

	for i := 0; i < len(elements)/2; i++ {
		elements[i], elements[last-i] = elements[last-i], elements[i]
	}
}
