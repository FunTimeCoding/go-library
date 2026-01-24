package example

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/tool/common"
)

func DeleteSilence() {
	argument.ParseBind()
	common.Alertmanager().DeleteSilence(
		argument.RequiredPositional(0, "SILENCE_IDENTIFIER"),
	)
}
