package example

import "github.com/funtimecoding/go-library/pkg/prometheus/alertmanager"

func DeleteSilence() {
	alertmanager.NewEnvironment().DeleteSilence(
		"81d9935b-0990-4761-9615-672e02d70340",
	)
}
