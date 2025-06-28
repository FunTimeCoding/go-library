package job

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/gitlab"
	"github.com/funtimecoding/go-library/pkg/gitlab/check/job/option"
	"github.com/funtimecoding/go-library/pkg/gitlab/constant"
)

func Check(o *option.Job) {
	g := gitlab.NewEnvironment()
	relevant := g.FailedJobs(o.Verbose)

	if o.Notation {
		printNotation(relevant, o)

		return
	}

	if len(relevant) == 0 {
		fmt.Println("No failed jobs found.")

		return
	}

	f := constant.Format

	for _, j := range relevant {
		fmt.Printf("Project: %s\n", j.Project.Format(f))
		fmt.Printf("  Job: %s\n", j.Format(f))
	}
}
