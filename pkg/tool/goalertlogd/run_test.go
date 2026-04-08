package goalertlogd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/mock_client"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/client"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/poller"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/route"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/server"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/store"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/web"
	"io"
	"net/http"
	"path/filepath"
	"testing"
	"time"
)

func TestRunLifecycle(t *testing.T) {
	path := filepath.Join(t.TempDir(), "test.db")
	s := store.New(path)
	defer s.Close()
	c := mock_client.New()
	c.Add(
		alert.NewBasic(
			"fp1",
			"HighMemory",
			"critical",
			"Memory above 90%",
		),
	)
	c.Add(
		alert.NewBasic(
			"fp2",
			"DiskFull",
			"warning",
			"Disk usage above 85%",
		),
	)
	g := logger.New(context.Background())
	p := poller.New(c, s, g, 1*time.Minute, 30*24*time.Hour, nil)
	p.Poll()
	port := system.FindUnusedPort(19400)
	address := fmt.Sprintf(":%d", port)
	l := lifecycle.New(
		lifecycle.WithWorker(p),
		lifecycle.WithServerTimeout(
			address,
			func(m *http.ServeMux) {
				server.HandlerFromMux(route.New(s, p), m)
				web.NewServer(s, p).Mount(m)
			},
			0,
			0,
		),
	)
	l.Run()
	defer l.Stop()
	assert.Listen(t, port)
	base := fmt.Sprintf("http://localhost:%d", port)
	assert.HTTPStatus(t, fmt.Sprintf("%s/", base), http.StatusOK)
	assert.HTTPStatus(t, fmt.Sprintf("%s/recent", base), http.StatusOK)
	assert.HTTPStatus(
		t,
		fmt.Sprintf("%s/alerts?name=HighMemory", base),
		http.StatusOK,
	)
	assert.HTTPStatus(t, fmt.Sprintf("%s/api/v1/status", base), http.StatusOK)
	assert.HTTPStatus(
		t,
		fmt.Sprintf("%s/api/v1/alerts?name=HighMemory", base),
		http.StatusOK,
	)
	assert.HTTPStatus(
		t,
		fmt.Sprintf("%s/api/v1/alerts/recent", base),
		http.StatusOK,
	)
	assert.HTTPStatus(
		t,
		fmt.Sprintf("%s/api/v1/alerts/top", base),
		http.StatusOK,
	)
	assert.HTTPStatus(
		t,
		fmt.Sprintf("%s/api/v1/alerts", base),
		http.StatusBadRequest,
	)
	status := getJSON[server.StatusResponse](
		t,
		fmt.Sprintf("%s/api/v1/status", base),
	)
	assert.Integer(t, 2, status.TotalRecords)
	assert.True(t, status.LastPoll != nil)
	alerts := getJSON[[]server.AlertsResponse](
		t,
		fmt.Sprintf("%s/api/v1/alerts?name=HighMemory", base),
	)
	assert.Count(t, 1, alerts)
	assert.String(t, "fp1", alerts[0].Fingerprint)
	assert.String(t, string(server.Firing), string(alerts[0].Status))
	recent := getJSON[[]server.AlertsResponse](
		t,
		fmt.Sprintf("%s/api/v1/alerts/recent", base),
	)
	assert.Count(t, 2, recent)
	top := getJSON[[]server.TopAlertsResponse](
		t,
		fmt.Sprintf("%s/api/v1/alerts/top", base),
	)
	assert.Count(t, 2, top)
	assert.Integer(t, 1, top[0].Count)
	c.Remove("fp1")
	p.Poll()
	alerts = getJSON[[]server.AlertsResponse](
		t,
		fmt.Sprintf("%s/api/v1/alerts?name=HighMemory", base),
	)
	assert.Count(t, 1, alerts)
	assert.String(t, string(server.Resolved), string(alerts[0].Status))
	assert.True(t, alerts[0].End != nil)
	status = getJSON[server.StatusResponse](
		t,
		fmt.Sprintf("%s/api/v1/status", base),
	)
	assert.Integer(t, 2, status.TotalRecords)
}

func TestGeneratedClient(t *testing.T) {
	path := filepath.Join(t.TempDir(), "test.db")
	s := store.New(path)
	defer s.Close()
	c := mock_client.New()
	c.Add(
		alert.NewBasic(
			"fp1",
			"HighMemory",
			"critical",
			"Memory above 90%",
		),
	)
	g := logger.New(context.Background())
	p := poller.New(c, s, g, 1*time.Minute, 30*24*time.Hour, nil)
	p.Poll()
	port := system.FindUnusedPort(19500)
	address := fmt.Sprintf(":%d", port)
	l := lifecycle.New(
		lifecycle.WithWorker(p),
		lifecycle.WithServerTimeout(
			address,
			func(m *http.ServeMux) {
				server.HandlerFromMux(route.New(s, p), m)
			},
			0,
			0,
		),
	)
	l.Run()
	defer l.Stop()
	assert.Listen(t, port)
	base := fmt.Sprintf("http://localhost:%d", port)
	cl, e := client.NewClientWithResponses(base)
	assert.FatalOnError(t, e)
	x := context.Background()
	status, f := cl.GetStatusWithResponse(x)
	assert.FatalOnError(t, f)
	assert.Integer(t, http.StatusOK, status.StatusCode())
	assert.NotNil(t, status.JSON200)
	assert.Integer(t, 1, status.JSON200.TotalRecords)
	assert.True(t, status.JSON200.LastPoll != nil)
	alerts, h := cl.GetAlertsWithResponse(
		x,
		&client.GetAlertsParams{Name: "HighMemory"},
	)
	assert.FatalOnError(t, h)
	assert.Integer(t, http.StatusOK, alerts.StatusCode())
	assert.NotNil(t, alerts.JSON200)
	assert.Count(t, 1, *alerts.JSON200)
	assert.String(t, "fp1", (*alerts.JSON200)[0].Fingerprint)
	assert.String(
		t,
		string(client.Firing),
		string((*alerts.JSON200)[0].Status),
	)
	recent, i := cl.GetRecentAlertsWithResponse(
		x,
		&client.GetRecentAlertsParams{},
	)
	assert.FatalOnError(t, i)
	assert.Integer(t, http.StatusOK, recent.StatusCode())
	assert.NotNil(t, recent.JSON200)
	assert.Count(t, 1, *recent.JSON200)
	topAlerts, j := cl.GetTopAlertsWithResponse(
		x,
		&client.GetTopAlertsParams{},
	)
	assert.FatalOnError(t, j)
	assert.Integer(t, http.StatusOK, topAlerts.StatusCode())
	assert.NotNil(t, topAlerts.JSON200)
	assert.Count(t, 1, *topAlerts.JSON200)
	assert.Integer(t, 1, (*topAlerts.JSON200)[0].Count)
}

func getJSON[T any](
	t *testing.T,
	locator string,
) T {
	t.Helper()
	r, e := http.Get(locator)
	assert.FatalOnError(t, e)
	defer errors.PanicClose(r.Body)
	body, f := io.ReadAll(r.Body)
	errors.PanicOnError(f)
	var result T
	errors.PanicOnError(json.Unmarshal(body, &result))

	return result
}
