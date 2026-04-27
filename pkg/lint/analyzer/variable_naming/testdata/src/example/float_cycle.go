package example

func FloatCycle() {
	p := 1.0 // want `variable p of type float64 should be named f`
	q := 2.0 // want `variable q of type float64 should be named l`
	r := 3.0 // want `variable r of type float64 should be named o`
	_ = p + q + r
}
