package example

import "os"

func Flagged() {
	f, _ := os.Open("test")
	defer f.Close() // want `use defer errors\.PanicClose\(f\) instead of defer f\.Close\(\)`
}
