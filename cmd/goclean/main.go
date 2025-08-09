package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/git"
	"github.com/funtimecoding/go-library/pkg/git/constant"
	"github.com/funtimecoding/go-library/pkg/git/remote/provider_map"
	"github.com/funtimecoding/go-library/pkg/github"
	"github.com/funtimecoding/go-library/pkg/github/run"
	hubTag "github.com/funtimecoding/go-library/pkg/github/tag"
	"github.com/funtimecoding/go-library/pkg/gitlab"
	gitlabConstant "github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"github.com/funtimecoding/go-library/pkg/gitlab/image"
	"github.com/funtimecoding/go-library/pkg/gitlab/packages"
	"github.com/funtimecoding/go-library/pkg/gitlab/pipeline"
	labTag "github.com/funtimecoding/go-library/pkg/gitlab/tag"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/funtimecoding/go-library/pkg/web/locator"
	"github.com/spf13/viper"
	"slices"
	"strings"
)

func main() {
	monitor.VerboseArgument()
	argument.ParseBind()
	verbose := viper.GetBool(argument.Verbose)

	gitlabLocator := web.ParseLocator(
		environment.Get(gitlabConstant.HostEnvironment),
	)
	m := provider_map.New()
	m.Add(gitlabLocator.Host, provider_map.GitLabProvider)

	if locator.HasDot(gitlabLocator.Host) {
		m.Add(
			locator.RemoveSubdomain(gitlabLocator.Host),
			provider_map.GitLabProvider,
		)
	}

	origin := git.Remote(system.WorkingDirectory(), m, constant.OriginRemote)

	if origin == nil {
		system.Exitf(
			1,
			"could not identify provider: %s\n",
			gitlabLocator.Host,
		)

		return
	}

	switch origin.Provider {
	case provider_map.GitHubProvider:
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
		c := github.NewEnvironment()
		tags := c.Tags(namespace, repository)
		latestTag := hubTag.Latest(tags)

		if latestTag == nil {
			system.Exitf(
				1,
				"could not find latest tag for %s/%s\n",
				namespace,
				repository,
			)

			return
		}

		for _, t := range tags {
			if t.Name == latestTag.Name {
				continue
			}

			fmt.Printf("Delete tag: %s\n", *t.Name)
			c.DeleteTag(namespace, repository, *t.Name)
		}

		git.Fetch()

		runs := c.ProjectRuns(namespace, repository)

		if false {
			latestRun := run.Latest(runs)
			fmt.Printf("Latest run: %s\n", latestRun.Name)
		}

		for _, r := range runs {
			if r.Status != run.Completed {
				continue
			}

			fmt.Printf("Delete run: %d\n", r.Identifier)
			c.DeleteRun(namespace, repository, r.Identifier)
		}
	case provider_map.GitLabProvider:
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
			gitlabLocator.Host,
			environment.Get(gitlabConstant.TokenEnvironment),
		)
		p := c.ProjectByName(namespace, repository)

		if verbose {
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

		if verbose {
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
						if verbose {
							fmt.Printf(
								"Skip pipeline (ref): %s %s\n",
								i.Ref,
								i.SHA,
							)
						}

						continue
					} else {
						if i.SHA == mainBranchHash {
							if verbose {
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
	case provider_map.UnknownProvider:
		// TODO: Consider deleting tags except latest locally and pushing them to the server
		fmt.Println("Unknown provider, nothing to clean")
	}
}
