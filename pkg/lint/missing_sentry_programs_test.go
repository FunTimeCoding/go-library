package lint

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/system/virtual_file_system"
	"testing"
)

func TestMissingSentryProgramsFlagged(t *testing.T) {
	v := virtual_file_system.New()
	v.Write("cmd/gotest/main.go", "package main\n")
	v.Write("pkg/tool/gotest/main.go", "package gotest\n\nfunc Main() {}\n")
	result := missingSentryPrograms(v)
	assert.Integer(t, 1, len(result))
	assert.String(t, "gotest", result[0])
}

func TestMissingSentryProgramsWithReporter(t *testing.T) {
	v := virtual_file_system.New()
	v.Write("cmd/gotest/main.go", "package main\n")
	v.Write(
		"pkg/tool/gotest/run.go",
		"package gotest\n\nfunc Run() {\n\tr := reporter.New(\"gotest\", c, \"\", o.Version)\n}\n",
	)
	result := missingSentryPrograms(v)
	assert.Integer(t, 0, len(result))
}

func TestMissingSentryProgramsSkipsExample(t *testing.T) {
	v := virtual_file_system.New()
	v.Write("cmd/example/main.go", "package main\n")
	v.Write("pkg/tool/example/main.go", "package example\n\nfunc Main() {}\n")
	result := missingSentryPrograms(v)
	assert.Integer(t, 0, len(result))
}

func TestMissingSentryProgramsSorted(t *testing.T) {
	v := virtual_file_system.New()
	v.Write("cmd/gobravo/main.go", "package main\n")
	v.Write("pkg/tool/gobravo/main.go", "package gobravo\n")
	v.Write("cmd/goalfa/main.go", "package main\n")
	v.Write("pkg/tool/goalfa/main.go", "package goalfa\n")
	result := missingSentryPrograms(v)
	assert.Integer(t, 2, len(result))
	assert.String(t, "goalfa", result[0])
	assert.String(t, "gobravo", result[1])
}
