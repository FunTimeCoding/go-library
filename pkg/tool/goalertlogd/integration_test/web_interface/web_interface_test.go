//go:build local

package web_interface

import (
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/integration_test/web_interface_tester"
	"net/http"
	"testing"
)

func TestWebInterface(t *testing.T) {
	o := web_interface_tester.New(t)
	defer o.Close()
	o.AssertStatus(constant.DashboardPath, http.StatusOK)
	o.AssertStatus(constant.RecentPath, http.StatusOK)
	o.AssertStatus("/alerts?name=HighMemory", http.StatusOK)
}
