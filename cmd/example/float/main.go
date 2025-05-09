package main

import "fmt"

func main() {
	f := 1.52
	fmt.Printf("As f: %f\n", f)
	fmt.Printf("As f.1: %.1f\n", f)
	fmt.Printf("As f.2: %.2f\n", f)
	fmt.Printf("As v: %v\n", f)
	fmt.Printf("As +v: %+v\n", f)
}
