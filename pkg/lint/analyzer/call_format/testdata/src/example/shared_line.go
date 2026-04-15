package example

func sharedLineArgs() {
	withStruct(
		"name", Options{ // want `each argument should be on its own line`
			Value: 1,
		},
	)
}
