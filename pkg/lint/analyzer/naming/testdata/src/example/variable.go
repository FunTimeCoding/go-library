package example

var (
	url             = "test"   // want `use "l" or "locator" instead of "url" in url`
	locator         = "test"   // ok: uses the suggestion
	dirName         = "test"   // want `use "directory" instead of "dir" in dirName`
	directory       = "test"   // ok: contains the suggestion
	DirSomething    = "test"   // want `use "directory" instead of "dir" in DirSomething`
	OutputDirectory = ""       // ok: segment is "directory"
	ctx             = ""       // want `use "x" instead of "ctx" in ctx`
	msgText         = ""       // want `use "message" instead of "msg" in msgText`
	paramCount      = 0        // want `use "parameter" instead of "param" in paramCount`
	parameterCount  = 0        // ok: contains the suggestion
	mcpServer       = ""       // want `use "model_context" instead of "mcp" in mcpServer`
	clean           = ""       // ok: no blacklisted segment
	req             = ""       // want `use "r" or "request" instead of "req" in req`
	request         = ""       // ok: uses the suggestion
	docPath         = ""       // want `use "document" instead of "doc" in docPath`
	document        = ""       // ok: uses the suggestion
	cfg             = ""       // want `use "c" or "configuration" instead of "cfg" in cfg`
	llm             = ""       // want `use "m" or "model" instead of "llm" in llm`
	tmp             = ""       // want `use "t" instead of "tmp" in tmp`
	loginHandler    = ""       // want `avoid "handler" in name, use a more specific term`
	userInfo        = ""       // want `avoid "info" in name, use a more specific term`
	userData        = ""       // want `avoid "data" in name, use a more specific term`
	posX            = 0        // want `use "position" instead of "pos" in posX`
	position        = 0        // ok: uses the suggestion
	bufData         = []byte{} // want `use "buffer" instead of "buf" in bufData`
	buffer          = []byte{} // ok: uses the suggestion
	ptrValue        = 0        // want `use "pointer" instead of "ptr" in ptrValue`
	addrStr         = ""       // want `use "address" instead of "addr" in addrStr`
	refCount        = 0        // want `use "reference" instead of "ref" in refCount`
	navBar          = ""       // want `use "navigation" instead of "nav" in navBar`
	prevState       = ""       // want `use "past" instead of "prev" in prevState`
	past            = ""       // ok: uses the suggestion
	yamlData        = []byte{} // want `use "markup" instead of "yaml" in yamlData`
	xmlContent      = []byte{} // want `use "markup" instead of "xml" in xmlContent`
	htmlPage        = ""       // want `use "markup" instead of "html" in htmlPage`
	jsonPayload     = []byte{} // want `use "notation" instead of "json" in jsonPayload`
	declCount       = 0        // want `use "declaration" instead of "decl" in declCount`
)
