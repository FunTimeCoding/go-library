package status

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
	"github.com/funtimecoding/go-library/pkg/git/check/status/option"
	"github.com/funtimecoding/go-library/pkg/git/constant"
	"github.com/funtimecoding/go-library/pkg/monitor"
)

func Check(o *option.Status) {
	elements := monitor.OnlyConcerns(collect(o.Path, o.Depth), o.All)

	if o.Notation {
		printNotation(elements, o)

		return
	}

	f := constant.Format.Copy()

	if o.Verbose {
		f.Tag(tag.Changes)
	}

	for _, r := range elements {
		fmt.Println(r.Format(f))
	}

	if len(elements) == 0 {
		monitor.NoRelevant(Plural)
	}
}
