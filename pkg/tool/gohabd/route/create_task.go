package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/convert"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/habitica"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (h *Router) CreateTask(
	w http.ResponseWriter,
	q *http.Request,
) {
	var b server.CreateTaskJSONRequestBody
	errors.PanicOnError(json.NewDecoder(q.Body).Decode(&b))
	body := habitica.CreateTaskBody{
		Type: string(b.Type),
		Text: b.Text,
	}

	if b.Notes != nil {
		body.Notes = *b.Notes
	}

	web.ObjectHeader(w)
	w.WriteHeader(http.StatusCreated)
	web.Encode(w, convert.Task(h.habitica.CreateTask(body)))
}
