package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/convert"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (r *Router) UpdatePage(
	w http.ResponseWriter,
	e *http.Request,
	identifier string,
) {
	var body server.UpdatePageJSONRequestBody
	errors.PanicOnError(json.NewDecoder(e.Body).Decode(&body))
	message := ""

	if body.Message != nil {
		message = *body.Message
	}

	web.EncodeNotation(
		w,
		convert.ConfluencePageDetail(
			r.confluence.UpdatePage(
				identifier,
				body.Title,
				body.Body,
				message,
			),
		),
	)
}
