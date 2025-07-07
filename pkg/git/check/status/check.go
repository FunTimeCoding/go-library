package status

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
	"github.com/funtimecoding/go-library/pkg/git/check/status/option"
	"github.com/funtimecoding/go-library/pkg/git/constant"
	"github.com/funtimecoding/go-library/pkg/git/repository"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func Check(o *option.Status) {
	elements := repository.Filter(
		monitor.OnlyConcerns(collect(o.Path, o.Depth), o.All),
		environment.GetSlice(RepositoryExcludeEnvironment),
		o.All,
	)

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
