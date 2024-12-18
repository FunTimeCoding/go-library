package postgres

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "psql", Command)
	assert.String(t, "--username", UserArgument)
	assert.String(t, "--command", CommandArgument)
	assert.String(t, "--file", FileArgument)
	assert.String(t, "--echo-all", EchoAllFlag)
	assert.String(t, "pg_dump", DumpCommand)
}
