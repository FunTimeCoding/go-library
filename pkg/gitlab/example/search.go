package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/gitlab"
	"github.com/funtimecoding/go-library/pkg/project"
)

func Search() {
	g := gitlab.NewEnvironment()
	// Free version search is limited

	for _, p := range g.SearchProject("") {
		fmt.Printf("Project: %s\n", p.Raw.NameWithNamespace)
	}

	for _, b := range g.SearchBlob(
		fmt.Sprintf("filename:%s", project.GitLabFile),
	) {
		fmt.Printf("Blob: %+v\n", b)
	}
}
