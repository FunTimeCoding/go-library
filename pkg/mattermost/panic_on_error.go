package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/mattermost/mattermost-server/v6/model"
	"log"
	"net/http"
	"slices"
)

func panicOnError(
	e error,
	r *model.Response,
) {
	errors.PanicOnError(e)

	if !slices.Contains(
		[]int{http.StatusOK, http.StatusCreated},
		r.StatusCode,
	) {
		log.Panicf("unexpected response code: %+v\n", r)
	}
}
