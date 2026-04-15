package example

func sharedLineArgs() {
	withStruct(
		"name", Options{ // want `each argument should be on its own line`
			Value: 1,
		},
	)
}

func firstArgOnParenLine() {
	withStruct("name", // want `each argument should be on its own line`
		Options{
			Value: 1,
		},
	)
}
