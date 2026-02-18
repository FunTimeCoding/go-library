package lifecycle

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system"
	"net/http"
	"testing"
)

type worker struct {
	tag string
	log *[]string
}

func (w *worker) Start() {
	*w.log = append(*w.log, fmt.Sprintf("start:%s", w.tag))
}

func (w *worker) Stop() {
	*w.log = append(*w.log, fmt.Sprintf("stop:%s", w.tag))
}

func TestRunWorkerStartOrder(t *testing.T) {
	var log []string
	l := New(
		WithWorker(&worker{log: &log, tag: "a"}),
		WithWorker(&worker{log: &log, tag: "b"}),
		WithWorker(&worker{log: &log, tag: "c"}),
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
		WithWorker(&worker{log: &log, tag: "a"}),
		WithWorker(&worker{log: &log, tag: "b"}),
		WithWorker(&worker{log: &log, tag: "c"}),
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
	port := system.FindUnusedPort(19000)
	address := fmt.Sprintf(":%d", port)
	l := New(
		WithServer(
			address,
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
	assert.Listen(t, port)
	assert.HTTPStatus(
		t,
		fmt.Sprintf("http://localhost:%d/health", port),
		http.StatusOK,
	)
}

func TestStopServerShutsDown(t *testing.T) {
	port := system.FindUnusedPort(19100)
	address := fmt.Sprintf(":%d", port)
	l := New(
		WithServer(
			address,
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
	assert.Listen(t, port)
	l.Stop()
	assert.NotListen(t, port)
}

func TestRunServerMiddleware(t *testing.T) {
	port := system.FindUnusedPort(19300)
	address := fmt.Sprintf(":%d", port)
	l := New(
		WithServerMiddleware(
			address,
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
	assert.Listen(t, port)
	r, e := http.Get(fmt.Sprintf("http://localhost:%d/health", port))

	if e != nil {
		t.Fatal(e)
	}

	defer errors.PanicClose(r.Body)
	assert.Integer(t, http.StatusOK, r.StatusCode)
	assert.String(t, "middleware", r.Header.Get("X-Test"))
}

func TestRunMixedOrder(t *testing.T) {
	var log []string
	port := system.FindUnusedPort(19200)
	address := fmt.Sprintf(":%d", port)
	l := New(
		WithWorker(&worker{log: &log, tag: "a"}),
		WithServer(
			address,
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
		WithWorker(&worker{log: &log, tag: "b"}),
	)
	l.Run()
	assert.Count(t, 3, log)
	assert.String(t, "start:a", log[0])
	assert.String(t, "start:server", log[1])
	assert.String(t, "start:b", log[2])
	assert.Listen(t, port)
	log = nil
	l.Stop()
	assert.Count(t, 2, log)
	assert.String(t, "stop:b", log[0])
	assert.String(t, "stop:a", log[1])
}
