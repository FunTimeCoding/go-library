package record

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/telemetry/constant"
)

func New(
	tool string,
	surface string,
	actor string,
	outcome string,
) *Record {
	if !constant.ValidSurface(surface) {
		panic(fmt.Sprintf("invalid telemetry surface: %q", surface))
	}

	return &Record{
		Tool:    tool,
		Surface: surface,
		Actor:   actor,
		Outcome: outcome,
	}
}
