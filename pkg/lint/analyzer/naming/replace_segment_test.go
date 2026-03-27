package naming

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestReplaceSegmentCamelCase(t *testing.T) {
	assert.String(t, "directoryName", replaceSegment("dirName", "dir", "directory"))
}

func TestReplaceSegmentPascalCase(t *testing.T) {
	assert.String(t, "DirectorySomething", replaceSegment("DirSomething", "dir", "directory"))
}

func TestReplaceSegmentSingleWord(t *testing.T) {
	assert.String(t, "locator", replaceSegment("url", "url", "locator"))
}

func TestReplaceSegmentSingleWordExported(t *testing.T) {
	assert.String(t, "Locator", replaceSegment("Url", "url", "locator"))
}

func TestReplaceSegmentMiddle(t *testing.T) {
	assert.String(t, "fooLocator", replaceSegment("fooUrl", "url", "locator"))
}

func TestReplaceSegmentMiddleExported(t *testing.T) {
	assert.String(t, "FooLocator", replaceSegment("FooUrl", "url", "locator"))
}

func TestReplaceSegmentSnakeCase(t *testing.T) {
	assert.String(t, "foo_directory", replaceSegment("foo_dir", "dir", "directory"))
}

func TestReplaceSegmentMultiWordCamel(t *testing.T) {
	assert.String(t, "modelContextServer", replaceSegment("mcpServer", "mcp", "model_context"))
}

func TestReplaceSegmentMultiWordPascal(t *testing.T) {
	assert.String(t, "ModelContextServer", replaceSegment("McpServer", "mcp", "model_context"))
}

func TestReplaceSegmentMultiWordSnake(t *testing.T) {
	assert.String(t, "model_context_server", replaceSegment("mcp_server", "mcp", "model_context"))
}

func TestReplaceSegmentSingleChar(t *testing.T) {
	assert.String(t, "xWrapper", replaceSegment("ctxWrapper", "ctx", "x"))
}

func TestReplaceSegmentSingleCharExported(t *testing.T) {
	assert.String(t, "XWrapper", replaceSegment("CtxWrapper", "ctx", "x"))
}

func TestReplaceSegmentSingleCharAlone(t *testing.T) {
	assert.String(t, "x", replaceSegment("ctx", "ctx", "x"))
}

func TestReplaceSegmentNoMatch(t *testing.T) {
	assert.String(t, "fooBar", replaceSegment("fooBar", "baz", "qux"))
}
