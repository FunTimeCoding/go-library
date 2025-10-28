package server

import (
	"context"
	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/errors"
	"net/http/httptest"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	t.Parallel()

	s := httptest.NewServer(Server{Logf: t.Logf})
	defer s.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	c, _, err := websocket.Dial(
		ctx,
		s.URL,
		&websocket.DialOptions{Subprotocols: []string{"echo"}},
	)
	assert.FatalOnError(t, err)
	defer func() {
		errors.LogOnError(c.CloseNow())
	}()

	for i := 0; i < 5; i++ {
		assert.FatalOnError(t, wsjson.Write(ctx, c, map[string]int{"i": i}))

		v := map[string]int{}
		assert.FatalOnError(t, wsjson.Read(ctx, c, &v))

		if v["i"] != i {
			t.Fatalf("expect %v but got %v", i, v)
		}
	}

	errors.LogOnError(c.Close(websocket.StatusNormalClosure, ""))
}
