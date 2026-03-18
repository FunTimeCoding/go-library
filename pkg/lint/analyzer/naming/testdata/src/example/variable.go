package example

var url = "test"          // want `use "locator" instead of "url" in url`
var locator = "test"      // ok: uses the suggestion
var dirName = "test"      // want `use "directory" instead of "dir" in dirName`
var directory = "test"    // ok: contains the suggestion
var DirSomething = "test" // want `use "directory" instead of "dir" in DirSomething`
var OutputDirectory = ""  // ok: segment is "directory"
var ctx = ""              // want `use "x" instead of "ctx" in ctx`
var msgText = ""          // want `use "message" instead of "msg" in msgText`
var paramCount = 0        // want `use "parameter" instead of "param" in paramCount`
var parameterCount = 0    // ok: contains the suggestion
var mcpServer = ""        // want `use "model_context" instead of "mcp" in mcpServer`
var clean = ""            // ok: no blacklisted segment
