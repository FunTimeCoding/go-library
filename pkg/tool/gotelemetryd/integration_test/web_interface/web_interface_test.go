//go:build local

package web_interface

import (
	"github.com/funtimecoding/go-library/pkg/tool/gotelemetryd/integration_test/web_interface_tester"
	"net/http"
	"testing"
)

func TestWebInterface(t *testing.T) {
	o := web_interface_tester.New(t)
	defer o.Close()
	o.AssertStatus("/", http.StatusOK)
	o.AssertStatus("/events", http.StatusOK)
}
