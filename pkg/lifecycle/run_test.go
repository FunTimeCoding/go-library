package lifecycle

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/lifecycle/mock_worker"
	"github.com/funtimecoding/go-library/pkg/system"
	"net/http"
	"testing"
)

func TestRunWorkerStartOrder(t *testing.T) {
	var log []string
	l := New(
		WithWorker(mock_worker.New("a", &log)),
		WithWorker(mock_worker.New("b", &log)),
		WithWorker(mock_worker.New("c", &log)),
	)
	l.Run()
	assert.Count(t, 3, log)
	assert.String(t, "start:a", log[0])
	assert.String(t, "start:b", log[1])
	assert.String(t, "start:c", log[2])
}

func TestStopWorkerReverseOrder(t *testing.T) {
	var log []string
	l := New(
		WithWorker(mock_worker.New("a", &log)),
		WithWorker(mock_worker.New("b", &log)),
		WithWorker(mock_worker.New("c", &log)),
	)
	l.Run()
	log = nil
	l.Stop()
	assert.Count(t, 3, log)
	assert.String(t, "stop:c", log[0])
	assert.String(t, "stop:b", log[1])
	assert.String(t, "stop:a", log[2])
}

func TestRunEmpty(t *testing.T) {
	l := New()
	l.Run()
}

func TestStopEmpty(t *testing.T) {
	l := New()
	l.Stop()
}

func TestRunServerResponds(t *testing.T) {
	p, n := system.ClaimPort()
	l := New(
		WithListener(
			n,
			func(m *http.ServeMux) {
				m.HandleFunc(
					"/health",
					func(
						w http.ResponseWriter,
						_ *http.Request,
					) {
						w.WriteHeader(http.StatusOK)
					},
				)
			},
		),
	)
	l.Run()
	defer l.Stop()
	assert.Listen(t, p)
	assert.HTTPStatus(
		t,
		fmt.Sprintf("http://localhost:%d/health", p),
		http.StatusOK,
	)
}

func TestStopServerShutsDown(t *testing.T) {
	p, n := system.ClaimPort()
	l := New(
		WithListener(
			n,
			func(m *http.ServeMux) {
				m.HandleFunc(
					"/health",
					func(
						w http.ResponseWriter,
						_ *http.Request,
					) {
						w.WriteHeader(http.StatusOK)
					},
				)
			},
		),
	)
	l.Run()
	assert.Listen(t, p)
	l.Stop()
	assert.NotListen(t, p)
}

func TestRunServerMiddleware(t *testing.T) {
	p := system.FindUnusedPort(19300)
	l := New(
		WithServerMiddleware(
			fmt.Sprintf(":%d", p),
			func(m *http.ServeMux) {
				m.HandleFunc(
					"/health",
					func(
						w http.ResponseWriter,
						_ *http.Request,
					) {
						w.WriteHeader(http.StatusOK)
					},
				)
			},
			func(h http.Handler) http.Handler {
				return http.HandlerFunc(
					func(
						w http.ResponseWriter,
						r *http.Request,
					) {
						w.Header().Set("X-Test", "middleware")
						h.ServeHTTP(w, r)
					},
				)
			},
		),
	)
	l.Run()
	defer l.Stop()
	assert.Listen(t, p)
	r, e := http.Get(fmt.Sprintf("http://localhost:%d/health", p))
	assert.FatalOnError(t, e)
	defer errors.PanicClose(r.Body)
	assert.Integer(t, http.StatusOK, r.StatusCode)
	assert.String(t, "middleware", r.Header.Get("X-Test"))
}

func TestRunMixedOrder(t *testing.T) {
	var log []string
	p, n := system.ClaimPort()
	l := New(
		WithWorker(mock_worker.New("a", &log)),
		WithListener(
			n,
			func(m *http.ServeMux) {
				log = append(log, "start:server")
				m.HandleFunc(
					"/health",
					func(
						w http.ResponseWriter,
						_ *http.Request,
					) {
						w.WriteHeader(http.StatusOK)
					},
				)
			},
		),
		WithWorker(mock_worker.New("b", &log)),
	)
	l.Run()
	assert.Count(t, 3, log)
	assert.String(t, "start:a", log[0])
	assert.String(t, "start:server", log[1])
	assert.String(t, "start:b", log[2])
	assert.Listen(t, p)
	log = nil
	l.Stop()
	assert.Count(t, 2, log)
	assert.String(t, "stop:b", log[0])
	assert.String(t, "stop:a", log[1])
}
