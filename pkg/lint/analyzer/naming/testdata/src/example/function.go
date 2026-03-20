package example

func FixtureUrl() {} // want `use "locator" instead of "url" in FixtureUrl`

func FixtureLocator() {} // ok

func HandleReq() {} // want `use "request" instead of "req" in HandleReq`

func HandleRequest() {} // ok: uses the suggestion

func HandleConfig() {} // want `use "configuration" instead of "config" in HandleConfig`

func HandleHandler() {} // want `avoid "handler" in name, use a more specific term`

func function(
	ctx string, // want `use "x" instead of "ctx" in ctx`
	msg string, // want `use "m" or "message" instead of "msg" in msg`
) {
	_ = ctx
	_ = msg
}
