package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/mattermost/mattermost/server/public/model"
	"log"
	"slices"
)

func panicOnError(
	e error,
	r *model.Response,
) {
	errors.PanicOnError(e)

	if !slices.Contains(constant.OkayStatusCodes, r.StatusCode) {
		log.Panicf("unexpected response code: %+v\n", r)
	}
}
