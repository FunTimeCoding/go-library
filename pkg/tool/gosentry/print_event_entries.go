package gosentry

import (
	"github.com/atlassian/go-sentry-api"
	"github.com/atlassian/go-sentry-api/datatype"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func printEventEntries(v sentry.Event) {
	for _, n := range v.Entries {
		kind, data, e := n.GetInterface()
		errors.PanicOnError(e)

		switch kind {
		case "exception":
			printException(data.(*datatype.Exception))
		case "message":
			printMessage(data.(*datatype.Message))
		}
	}
}
