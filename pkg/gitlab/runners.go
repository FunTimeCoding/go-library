package gitlab

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"github.com/funtimecoding/go-library/pkg/gitlab/runner"
	"gitlab.com/gitlab-org/api/client-go"
	"slices"
)

func (c *Client) Runners(all bool) []*runner.Runner {
	var result []*gitlab.Runner
	var number int

	for {
		var page []*gitlab.Runner
		var response *gitlab.Response
		options := &gitlab.ListRunnersOptions{
			ListOptions: gitlab.ListOptions{
				PerPage: constant.PerPage100,
				Page:    number,
			},
		}

		if all {
			page, response = c.allRunners(options)
		} else {
			page, response = c.ownerRunners(options)
		}

		if false {
			// Paging does not make sense, duplicates are returned
			headerKeys := []string{
				"X-Next-Page",
				"X-Page",
				"X-Prev-Page",
				"X-Total",
				"X-Total-Pages",
				"X-Per-Page",
			}

			for k, v := range response.Header {
				if !slices.Contains(headerKeys, k) {
					continue
				}

				fmt.Printf("  %d Header %s: %v\n", number, k, v)
			}
		}

		result = append(result, page...)

		if len(page) < constant.PerPage100 {
			break
		}

		number++
	}

	return c.enrichRunners(runner.NewSlice(runner.Deduplicate(result)))
}
