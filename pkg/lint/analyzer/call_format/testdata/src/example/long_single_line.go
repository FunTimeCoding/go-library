package example

func longSingleLine() {
	twoArgs("something-long-enough", "to-push-this-line-well-past-the-eighty-character-column-limit-easily") // want `call exceeds 80 characters`
}
