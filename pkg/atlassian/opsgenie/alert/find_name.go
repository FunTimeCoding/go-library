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

	if a.RawList != nil {
		// details and description not available
		message = a.RawList.Message
	}

	if a.RawDetail != nil {
		message = a.RawDetail.Message
		details = a.RawDetail.Details
		description = a.RawDetail.Description
	}

	if n, okay := details[constant.AlertnameLabel]; okay {
		return a.shortenAlert(n)
	}

	result = a.shortenAlert(message)

	if result == UnknownName {
		result = a.nameFromDescription(description)
	}

	if result == UnknownName {
		if a.RawList != nil {
			fmt.Printf("Unknown name (simple): %+v\n", a.RawList)
		}

		if a.RawDetail != nil {
			fmt.Printf("Unknown name (full): %+v\n", a.RawDetail)
		}
	}

	return result
}
