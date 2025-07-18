package job

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/gitlab/check/job/option"
	"github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"github.com/funtimecoding/go-library/pkg/monitor"
)

func Check(o *option.Job) {
	elements := collect(o)

	if o.Notation {
		printNotation(elements, o)

		return
	}

	f := constant.CheckFormat

	for _, e := range elements {
		fmt.Println(e.Format(f))
	}

	if len(elements) == 0 {
		monitor.NoRelevant(Plural)
	}
}
