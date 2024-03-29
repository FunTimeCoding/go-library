package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/git"
	"github.com/funtimecoding/go-library/pkg/git/remote"
	"github.com/funtimecoding/go-library/pkg/git/remote/provider_map"
	"github.com/funtimecoding/go-library/pkg/github"
	"github.com/funtimecoding/go-library/pkg/gitlab"
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
			latest := gitlab.LatestPipeline(pipelines)

			for _, pipeline := range pipelines {
				if pipeline.Ref == latest.Ref {
					continue
				}

				fmt.Printf("Delete pipeline: %s\n", pipeline.Ref)
				c.DeletePipeline(p.Identifier, pipeline.ID)
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

			latest := gitlab.LatestImage(images)

			for _, image := range images {
				if strings.HasSuffix(image.Path, ":latest") {
					continue
				}

				if image.Path == latest.Path {
					continue
				}

				fmt.Printf("Delete image: %s\n", image.Name)
				c.DeleteImage(p.Identifier, element.ID, image.Name)
			}
		}

		packages := c.Packages(p.Identifier, false)

		if len(packages) > 0 {
			latest := gitlab.LatestPackages(packages)

			for _, element := range packages {
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
			latest := gitlab.LatestTag(tags)

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
