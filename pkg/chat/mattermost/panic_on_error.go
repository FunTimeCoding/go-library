package mattermost

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/mattermost/mattermost/server/public/model"
	"log"
	"slices"
)

func panicOnError(
	r *model.Response,
	e error,
) {
	if e != nil && r != nil {
		fmt.Printf("Status: %d\n", r.StatusCode)
		web.PrintHeader(r.Header)
		fmt.Printf("Response: %+v\n", r)
	}

	errors.PanicOnError(e)

	if !slices.Contains(constant.OkayStatusCodes, r.StatusCode) {
		log.Panicf("unexpected response code: %+v\n", r)
	}
}
