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
	"io"
	"net/http"
	"testing"
	"time"
)

func TestRunLifecycle(t *testing.T) {
	path := t.TempDir() + "/test.db"
	s := store.New(path)
	defer s.Close()
	c := mock_client.New()
	c.Add(
		&alert.Alert{
			Fingerprint: "fp1",
			Name:        "HighMemory",
			Severity:    "critical",
			Summary:     "Memory above 90%",
		},
	)
	c.Add(
		&alert.Alert{
			Fingerprint: "fp2",
			Name:        "DiskFull",
			Severity:    "warning",
			Summary:     "Disk usage above 85%",
		},
	)
	g := logger.New(context.Background())
	p := poller.New(c, s, g, 1*time.Minute, 30*24*time.Hour)
	p.Poll()
	port := system.FindUnusedPort(19400)
	address := fmt.Sprintf(":%d", port)
	l := lifecycle.New(
		lifecycle.WithWorker(p),
		lifecycle.WithServer(
			address,
			func(m *http.ServeMux) {
				server.HandlerFromMux(route.New(s, p), m)
			},
		),
	)
	l.Run()
	defer l.Stop()
	assert.Listen(t, port)
	base := fmt.Sprintf("http://localhost:%d", port)
	assert.HTTPStatus(t, base+"/api/v1/status", http.StatusOK)
	assert.HTTPStatus(
		t,
		base+"/api/v1/alerts?name=HighMemory",
		http.StatusOK,
	)
	assert.HTTPStatus(
		t,
		base+"/api/v1/alerts/recent",
		http.StatusOK,
	)
	assert.HTTPStatus(
		t,
		base+"/api/v1/alerts",
		http.StatusBadRequest,
	)
	status := getJSON[server.StatusResponse](t, base+"/api/v1/status")
	assert.Integer(t, 2, status.TotalRecords)
	assert.True(t, status.LastPoll != nil)
	alerts := getJSON[[]server.AlertsResponse](
		t,
		base+"/api/v1/alerts?name=HighMemory",
	)
	assert.Count(t, 1, alerts)
	assert.String(t, "fp1", alerts[0].Fingerprint)
	assert.String(t, string(server.Firing), string(alerts[0].Status))
	recent := getJSON[[]server.AlertsResponse](
		t,
		base+"/api/v1/alerts/recent",
	)
	assert.Count(t, 2, recent)
	c.Remove("fp1")
	p.Poll()
	alerts = getJSON[[]server.AlertsResponse](
		t,
		base+"/api/v1/alerts?name=HighMemory",
	)
	assert.Count(t, 1, alerts)
	assert.String(t, string(server.Resolved), string(alerts[0].Status))
	assert.True(t, alerts[0].End != nil)
	status = getJSON[server.StatusResponse](t, base+"/api/v1/status")
	assert.Integer(t, 2, status.TotalRecords)
}

func TestGeneratedClient(t *testing.T) {
	path := t.TempDir() + "/test.db"
	s := store.New(path)
	defer s.Close()
	c := mock_client.New()
	c.Add(
		&alert.Alert{
			Fingerprint: "fp1",
			Name:        "HighMemory",
			Severity:    "critical",
			Summary:     "Memory above 90%",
		},
	)
	g := logger.New(context.Background())
	p := poller.New(c, s, g, 1*time.Minute, 30*24*time.Hour)
	p.Poll()
	port := system.FindUnusedPort(19500)
	address := fmt.Sprintf(":%d", port)
	l := lifecycle.New(
		lifecycle.WithWorker(p),
		lifecycle.WithServer(
			address,
			func(m *http.ServeMux) {
				server.HandlerFromMux(route.New(s, p), m)
			},
		),
	)
	l.Run()
	defer l.Stop()
	assert.Listen(t, port)
	base := fmt.Sprintf("http://localhost:%d", port)
	cl, e := client.NewClientWithResponses(base)

	if e != nil {
		t.Fatal(e)
	}

	ctx := context.Background()
	status, e := cl.GetStatusWithResponse(ctx)

	if e != nil {
		t.Fatal(e)
	}

	assert.Integer(t, http.StatusOK, status.StatusCode())
	assert.NotNil(t, status.JSON200)
	assert.Integer(t, 1, status.JSON200.TotalRecords)
	assert.True(t, status.JSON200.LastPoll != nil)
	alerts, e := cl.GetAlertsWithResponse(
		ctx,
		&client.GetAlertsParams{Name: "HighMemory"},
	)

	if e != nil {
		t.Fatal(e)
	}

	assert.Integer(t, http.StatusOK, alerts.StatusCode())
	assert.NotNil(t, alerts.JSON200)
	assert.Count(t, 1, *alerts.JSON200)
	assert.String(t, "fp1", (*alerts.JSON200)[0].Fingerprint)
	assert.String(
		t,
		string(client.Firing),
		string((*alerts.JSON200)[0].Status),
	)
	recent, e := cl.GetRecentAlertsWithResponse(
		ctx,
		&client.GetRecentAlertsParams{},
	)

	if e != nil {
		t.Fatal(e)
	}

	assert.Integer(t, http.StatusOK, recent.StatusCode())
	assert.NotNil(t, recent.JSON200)
	assert.Count(t, 1, *recent.JSON200)
}

func getJSON[T any](
	t *testing.T,
	locator string,
) T {
	t.Helper()
	r, e := http.Get(locator)

	if e != nil {
		t.Fatal(e)
	}

	defer errors.PanicClose(r.Body)
	body, e := io.ReadAll(r.Body)
	errors.PanicOnError(e)

	var result T
	errors.PanicOnError(json.Unmarshal(body, &result))

	return result
}
