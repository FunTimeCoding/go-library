package job

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/gitlab"
	"github.com/funtimecoding/go-library/pkg/project"
)

func Check() {
	// TODO: Implement for GitLab like GitHub
	g := gitlab.NewEnvironment()

	if true {
		for _, p := range g.ProjectsWithFile(project.GitLabFile) {
			fmt.Printf("Project: %s\n", p.NameWithNamespace)
		}
	}

	if false {
		// Free version search is limited

		for _, p := range g.SearchProject("") {
			fmt.Printf("Project: %s\n", p.NameWithNamespace)
		}

		for _, b := range g.SearchBlob(
			fmt.Sprintf("filename:%s", project.GitLabFile),
		) {
			fmt.Printf("Blob: %+v\n", b)
		}
	}
}
