package example

func MapCycle() {
	u := map[string]int{} // want `variable u of type map\[string\]int should be named m`
	v := map[string]int{} // want `variable v of type map\[string\]int should be named a`
	w := map[string]int{} // want `variable w of type map\[string\]int should be named p`
	_ = u
	_ = v
	_ = w
}
