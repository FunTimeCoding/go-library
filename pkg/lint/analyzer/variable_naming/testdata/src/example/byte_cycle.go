package example

func ByteCycle() {
	u := []byte("first")  // want `variable u of type \[\]byte should be named b`
	v := []byte("second") // want `variable v of type \[\]byte should be named y`
	w := []byte("third")  // want `variable w of type \[\]byte should be named t`
	_ = u
	_ = v
	_ = w
}
