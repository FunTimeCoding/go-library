package example

func shortCall() {
	twoArgs("alpha", "bravo")
}

func alreadySplit() {
	twoArgs(
		"alpha",
		"bravo",
	)
}

func splitWithStruct() {
	withStruct(
		"name",
		Options{
			Value: 1,
		},
	)
}

func singleArg() {
	oneArg("this is a very long string that would exceed eighty characters if there were more arguments")
}

func twoArgs(a, b string)    {}
func withStruct(a string, b Options) {}
func oneArg(a string)        {}

type Options struct {
	Value int
}
