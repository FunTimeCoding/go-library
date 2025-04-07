package alert

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
)

func (a *Alert) findName() string {
	var result string
	var details map[string]string
	var message string
	var description string

	if a.Raw != nil {
		// details and description not available
		message = a.Raw.Message
	}

	if a.RawResult != nil {
		message = a.RawResult.Message
		details = a.RawResult.Details
		description = a.RawResult.Description
	}

	if n, okay := details[constant.AlertnameLabel]; okay {
		return a.shortenAlert(n)
	}

	result = a.shortenAlert(message)

	if result == UnknownName {
		result = a.nameFromDescription(description)
	}

	if result == UnknownName {
		if a.Raw != nil {
			fmt.Printf("Unknown name (simple): %+v\n", a.Raw)
		}

		if a.RawResult != nil {
			fmt.Printf("Unknown name (full): %+v\n", a.RawResult)
		}
	}

	return result
}
