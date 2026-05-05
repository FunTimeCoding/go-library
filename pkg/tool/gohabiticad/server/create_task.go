package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/habitica/request"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/convert"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) CreateTask(
	w http.ResponseWriter,
	q *http.Request,
) {
	var b server.CreateTaskJSONRequestBody
	errors.PanicOnError(json.NewDecoder(q.Body).Decode(&b))
	body := request.CreateTaskBody{
		Type: string(b.Type),
		Text: b.Text,
	}

	if b.Notes != nil {
		body.Notes = *b.Notes
	}

	web.ObjectHeader(w)
	w.WriteHeader(http.StatusCreated)
	web.Encode(w, convert.Task(s.habitica.MustCreateTask(body)))
}
