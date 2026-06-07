package scan

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/system/virtual_file_system"
	"testing"
)

func TestServicesDetectsCapabilities(t *testing.T) {
	v := virtual_file_system.New()
	v.Write("pkg/tool/gotestd/server/router.go", "package server\n")
	v.Write("pkg/tool/gotestd/option/option.go", "package option\n")
	v.Write("pkg/tool/gotestd/run.go", "package gotestd\n")
	services := Services(v, "test")
	assert.Integer(t, 1, len(services))
	assert.String(t, "gotestd", services[0].Name)
	assert.Boolean(t, true, services[0].Server)
}

func TestServicesSkipsNonGoPrefix(t *testing.T) {
	v := virtual_file_system.New()
	v.Write("pkg/tool/other/server/router.go", "package server\n")
	services := Services(v, "test")
	assert.Integer(t, 0, len(services))
}

func TestServicesSkipsNoCapabilities(t *testing.T) {
	v := virtual_file_system.New()
	v.Write("pkg/tool/gotest/main.go", "package gotest\n")
	services := Services(v, "test")
	assert.Integer(t, 0, len(services))
}

func TestServicesSorted(t *testing.T) {
	v := virtual_file_system.New()
	v.Write("pkg/tool/gobravod/server/r.go", "package server\n")
	v.Write("pkg/tool/gobravod/option/o.go", "package option\n")
	v.Write("pkg/tool/gobravod/run.go", "package gobravod\n")
	v.Write("pkg/tool/goalfad/server/r.go", "package server\n")
	v.Write("pkg/tool/goalfad/option/o.go", "package option\n")
	v.Write("pkg/tool/goalfad/run.go", "package goalfad\n")
	services := Services(v, "test")
	assert.Integer(t, 2, len(services))
	assert.String(t, "goalfad", services[0].Name)
	assert.String(t, "gobravod", services[1].Name)
}
