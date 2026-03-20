package example

var url = "test"          // want `use "l" or "locator" instead of "url" in url`
var locator = "test"      // ok: uses the suggestion
var dirName = "test"      // want `use "d" or "directory" instead of "dir" in dirName`
var directory = "test"    // ok: contains the suggestion
var DirSomething = "test" // want `use "d" or "directory" instead of "dir" in DirSomething`
var OutputDirectory = ""  // ok: segment is "directory"
var ctx = ""              // want `use "x" instead of "ctx" in ctx`
var msgText = ""          // want `use "m" or "message" instead of "msg" in msgText`
var paramCount = 0        // want `use "parameter" instead of "param" in paramCount`
var parameterCount = 0    // ok: contains the suggestion
var mcpServer = ""        // want `use "c" or "model_context" instead of "mcp" in mcpServer`
var clean = ""            // ok: no blacklisted segment

var req = ""        // want `use "r" or "request" instead of "req" in req`
var request = ""    // ok: uses the suggestion
var docPath = ""    // want `use "d" or "document" instead of "doc" in docPath`
var document = ""   // ok: uses the suggestion
var cfg = ""        // want `use "c" or "configuration" instead of "cfg" in cfg`
var llm = ""        // want `use "m" or "model" instead of "llm" in llm`
var tmp = ""        // want `use "t" instead of "tmp" in tmp`

var loginHandler = "" // want `avoid "handler" in name, use a more specific term`
var userInfo = ""     // want `avoid "info" in name, use a more specific term`
var userData = ""     // want `avoid "data" in name, use a more specific term`
