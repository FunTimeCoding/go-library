package example

func function(
	ctx string, // want `use "x" instead of "ctx" in ctx`
	msg string, // want `use "m" or "message" instead of "msg" in msg`
) {
	_ = ctx
	_ = msg
}
