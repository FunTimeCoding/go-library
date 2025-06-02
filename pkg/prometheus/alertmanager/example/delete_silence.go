package example

import (
	"github.com/funtimecoding/go-library/internal"
	"github.com/funtimecoding/go-library/pkg/argument"
)

func DeleteSilence() {
	argument.ParseBind()
	internal.Alertmanager().DeleteSilence(
		argument.RequiredPositional(
			0,
			"SILENCE_IDENTIFIER",
			1,
		),
	)
}
