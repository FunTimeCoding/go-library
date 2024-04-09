package packages

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/gitlab/project"
	"github.com/xanzy/go-gitlab"
)

func Link(
	host string,
	project *project.Project,
	p *gitlab.Package,
	f *gitlab.PackageFile,
	verbose bool,
) string {
	if verbose {
		fmt.Printf("project: %+v\n", project)
		fmt.Printf("package: %+v\n", p)
		fmt.Printf("file: %+v\n", f)
	}

	result := fmt.Sprintf(
		"https://%s/api/v4/projects/%d/packages/%s/%s/%s/%s",
		host,
		project.Identifier,
		p.PackageType,
		p.Name,
		p.Version,
		f.FileName,
	)

	if verbose {
		fmt.Printf("link: %s\n", result)
	}

	return result
}
