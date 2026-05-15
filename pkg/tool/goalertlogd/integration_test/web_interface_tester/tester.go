package web_interface_tester

import (
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/integration_test/base"
	"testing"
)

type Tester struct {
	t      *testing.T
	server *base.Server
	base   string
}
