package example

func firstArgOnParenLine() {
	withStruct("name", // want `each argument should be on its own line`
		Options{
			Value: 1,
		},
	)
}
