package example

func FixtureUrl() {} // want `use "locator" instead of "url" in FixtureUrl`

func FixtureLocator() {} // ok

func function(
	ctx string, // want `use "x" instead of "ctx" in ctx`
	msg string, // want `use "message" instead of "msg" in msg`
) {
	_ = ctx
	_ = msg
}
