package panics

import "fmt"

func Print(
	f string,
	a ...any,
) {
	if len(a) > 0 {
		panic(fmt.Sprintf(f, a...))
	}

	panic(f)
}
