package job

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/gitlab"
	"github.com/funtimecoding/go-library/pkg/gitlab/check/job/option"
	"github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"github.com/funtimecoding/go-library/pkg/monitor"
)

func Check(o *option.Job) {
	g := gitlab.NewEnvironment()
	elements := g.FailedJobs(o.Verbose)

	if o.Notation {
		printNotation(elements, o)

		return
	}

	f := constant.Format

	for _, e := range elements {
		fmt.Printf("Project: %s\n", e.Project.Format(f))
		fmt.Printf("  Job: %s\n", e.Format(f))
	}

	if len(elements) == 0 {
		monitor.NoRelevant(Plural)
	}
}
