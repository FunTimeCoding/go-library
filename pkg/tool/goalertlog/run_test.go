package goalertlog

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
	"github.com/funtimecoding/go-library/pkg/tool/goalertlog/poller"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlog/route"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlog/store"
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
				m.HandleFunc("/api/v1/alerts", route.Alerts(s))
				m.HandleFunc("/api/v1/alerts/recent", route.Recent(s))
				m.HandleFunc("/api/v1/status", route.Status(s, p))
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
	status := getJSON[route.StatusResponse](t, base+"/api/v1/status")
	assert.Integer(t, 2, status.TotalRecords)
	assert.True(t, status.LastPoll != "")
	alerts := getJSON[[]route.AlertsResponse](
		t,
		base+"/api/v1/alerts?name=HighMemory",
	)
	assert.Count(t, 1, alerts)
	assert.String(t, "fp1", alerts[0].Fingerprint)
	assert.String(t, route.Firing, alerts[0].Status)
	recent := getJSON[[]route.AlertsResponse](t, base+"/api/v1/alerts/recent")
	assert.Count(t, 2, recent)
	c.Remove("fp1")
	p.Poll()
	alerts = getJSON[[]route.AlertsResponse](
		t,
		base+"/api/v1/alerts?name=HighMemory",
	)
	assert.Count(t, 1, alerts)
	assert.String(t, route.Resolved, alerts[0].Status)
	assert.True(t, alerts[0].End != "")
	status = getJSON[route.StatusResponse](t, base+"/api/v1/status")
	assert.Integer(t, 2, status.TotalRecords)
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
