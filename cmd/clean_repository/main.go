package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/git"
	"github.com/funtimecoding/go-library/pkg/git/remote"
	"github.com/funtimecoding/go-library/pkg/git/remote/provider_map"
	"github.com/funtimecoding/go-library/pkg/github"
	"github.com/funtimecoding/go-library/pkg/gitlab"
	"github.com/funtimecoding/go-library/pkg/gitlab/image"
	"github.com/funtimecoding/go-library/pkg/gitlab/packages"
	"github.com/funtimecoding/go-library/pkg/gitlab/pipeline"
	"github.com/funtimecoding/go-library/pkg/gitlab/tag"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/web"
	"os"
	"strings"
)

func main() {
	gitlabLocator := web.ParseLocator(environment.Get(gitlab.Host))
	m := provider_map.New()
	m.Add(gitlabLocator.Host, provider_map.GitLabProvider)
	var r *remote.Remote

	for _, element := range git.Remotes(system.WorkingDirectory(), m) {
		if element.Name == git.OriginRemote {
			r = element

			break
		}
	}

	if r == nil {
		fmt.Printf("Could not identify provider: %s", gitlabLocator.Host)

		os.Exit(1)
	}

	switch r.Provider {
	case provider_map.GitHubProvider:
		remoteLocator := git.ParseLocator(r.Locator)
		namespace, repository := git.ParseProject(remoteLocator.Path)
		c := github.New(environment.Get(github.Token))
		tags := c.Tags(namespace, repository)
		latest := github.LatestTag(tags)

		for _, element := range tags {
			if element.Name == latest.Name {
				continue
			}

			fmt.Printf("Delete tag: %s\n", *element.Name)
			c.DeleteTag(namespace, repository, *element.Name)
		}

		git.Fetch()
	case provider_map.GitLabProvider:
		remoteLocator := git.ParseLocator(r.Locator)
		namespace, repository := git.ParseProject(remoteLocator.Path)
		c := gitlab.New(gitlabLocator.Host, environment.Get(gitlab.Token))
		p := c.ProjectByName(namespace, repository)
		pipelines := c.Pipelines(p.Identifier)

		if len(pipelines) > 0 {
			latest := pipeline.Latest(pipelines)

			for _, element := range pipelines {
				if element.Ref == latest.Ref {
					continue
				}

				fmt.Printf("Delete pipeline: %s\n", element.Ref)
				c.DeletePipeline(p.Identifier, element.ID)
			}
		}

		for _, element := range c.RegistryRepositories(
			p.Identifier,
			false,
		) {
			images := c.Images(p.Identifier, element.ID)

			if len(images) == 0 {
				continue
			}

			latest := image.Latest(images)

			for _, i := range images {
				if strings.HasSuffix(i.Path, ":latest") {
					continue
				}

				if i.Path == latest.Path {
					continue
				}

				fmt.Printf("Delete image: %s\n", i.Name)
				c.DeleteImage(p.Identifier, element.ID, i.Name)
			}
		}

		projectPackages := c.Packages(p.Identifier, false)

		if len(projectPackages) > 0 {
			latest := packages.Latest(projectPackages)

			for _, element := range projectPackages {
				var isLatest bool

				for _, inner := range latest {
					if inner.Name == element.Name &&
						inner.Version == element.Version {
						isLatest = true
					}
				}

				if isLatest {
					continue
				}

				fmt.Printf(
					"Delete package: %s %s\n",
					element.Name,
					element.Version,
				)
				c.DeletePackage(p.Identifier, element.ID)
			}
		}

		tags := c.Tags(p.Identifier)

		if len(tags) > 0 {
			latest := tag.Latest(tags)

			for _, element := range tags {
				if element.Name == latest.Name {
					continue
				}

				fmt.Printf("Delete tag: %s\n", element.Name)
				c.DeleteTag(p.Identifier, element.Name)
			}

			git.Fetch()
		}
	case provider_map.UnknownProvider:
		// TODO: Consider deleting tags except latest locally and pushing them to the server
		fmt.Println("Unknown provider, nothing to clean")
	}
}
