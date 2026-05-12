package example

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/tool/common"
)

func DeleteSilence() {
	a := argument.NewSimple("delete-silence")
	a.ParseSimple()
	common.Alertmanager().DeleteSilence(
		a.RequiredPositional(0, "SILENCE_IDENTIFIER"),
	)
}
