package clean

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/git"
	"github.com/funtimecoding/go-library/pkg/git/constant"
	"github.com/funtimecoding/go-library/pkg/git/remote"
	"github.com/funtimecoding/go-library/pkg/gitlab"
	gitlabConstant "github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"github.com/funtimecoding/go-library/pkg/gitlab/image"
	"github.com/funtimecoding/go-library/pkg/gitlab/packages"
	"github.com/funtimecoding/go-library/pkg/gitlab/pipeline"
	labTag "github.com/funtimecoding/go-library/pkg/gitlab/tag"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/tool/goclean/clean/option"
	"slices"
	"strings"
)

func Lab(
	o *option.Clean,
	origin *remote.Remote,
) {
	remoteLocator := git.ParseLocator(origin.Locator)

	if remoteLocator == nil {
		system.Exitf(
			1,
			"could not parse remote locator: %s\n",
			origin.Locator,
		)

		return
	}

	namespace, repository := git.ParseProject(remoteLocator.Path)
	c := gitlab.New(
		o.GitLabHost,
		environment.Get(gitlabConstant.TokenEnvironment),
	)
	p := c.ProjectByName(namespace, repository)

	if o.Verbose {
		fmt.Printf(
			"Project: %d %s %s\n",
			p.Identifier,
			p.Namespace,
			p.Name,
		)
	}

	pipelines := c.Pipelines(p.Identifier)

	var mainBranchHash string

	for _, b := range c.Branches(p.Identifier) {
		if slices.Contains(constant.MainBranches, b.Name) {
			mainBranchHash = b.Raw.Commit.ID

			break
		}
	}

	if o.Verbose {
		fmt.Printf("Default branch: %s\n", p.Raw.DefaultBranch)

		for _, b := range c.Branches(p.Identifier) {
			fmt.Printf(
				"Branch: %s %s\n",
				b.Name,
				b.Raw.Commit.ID,
			)
		}

		fmt.Printf("Main branch hash: %s\n", mainBranchHash)
	}

	if len(pipelines) > 0 {
		latest := pipeline.Latest(pipelines)

		for _, i := range pipelines {
			if i.Ref == latest.Ref {
				if mainBranchHash == "" {
					if o.Verbose {
						fmt.Printf(
							"Skip pipeline (ref): %s %s\n",
							i.Ref,
							i.SHA,
						)
					}

					continue
				} else {
					if i.SHA == mainBranchHash {
						if o.Verbose {
							fmt.Printf(
								"Skip pipeline (sha): %s %s\n",
								i.Ref,
								i.SHA,
							)
						}

						continue
					}
				}
			}

			fmt.Printf("Delete pipeline: %s\n", i.Ref)
			c.DeletePipeline(p.Identifier, i.ID)
		}
	}

	for _, r := range c.RegistryRepositories(
		p.Identifier,
		false,
	) {
		// TODO: This still doesn't delete the latest, test now to see what happens
		images := c.Images(p.Identifier, r.ID)

		if len(images) == 0 {
			continue
		}

		latest := image.Latest(images)

		for _, i := range images {
			if strings.HasSuffix(i.Path, gitlabConstant.LatestSuffix) {
				continue
			}

			if i.Path == latest.Path {
				continue
			}

			fmt.Printf("Delete image: %s\n", i.Name)
			c.DeleteImage(p.Identifier, r.ID, i.Name)
		}
	}

	projectPackages := c.Packages(p.Identifier, false)

	if len(projectPackages) > 0 {
		latest := packages.Latest(projectPackages)

		for _, a := range projectPackages {
			var isLatest bool

			for _, inner := range latest {
				if inner.Name == a.Name &&
					inner.Version == a.Version {
					isLatest = true
				}
			}

			if isLatest {
				continue
			}

			fmt.Printf(
				"Delete package: %s %s\n",
				a.Name,
				a.Version,
			)
			c.DeletePackage(p.Identifier, a.ID)
		}
	}

	tags := c.Tags(p.Identifier)

	if len(tags) > 0 {
		latest := labTag.Latest(tags)

		for _, t := range tags {
			if t.Name == latest.Name {
				continue
			}

			fmt.Printf("Delete tag: %s\n", t.Name)
			c.DeleteTag(p.Identifier, t.Name)
		}

		git.Fetch()
	}
}
