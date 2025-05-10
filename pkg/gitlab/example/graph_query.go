package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/gitlab"
)

/*
Working

query {
  users(admins: true) { nodes { id username name publicEmail webUrl } }
}

query {
  runner(id: "gid://gitlab/Ci::Runner/1") {
    id description status runnerType
    managers { nodes { systemId ipAddress version revision } }
  }
}

query { project(fullPath: "gw2/wvw-schedule") { id } }

query { project(fullPath: "gw2/wvw-schedule") { id } }
*/

/*
Not working

query { runners(first: 5) { nodes { id description } } }
*/

type ProjectResult struct {
	Data struct {
		Project struct {
			ID string `json:"id"`
		} `json:"project"`
	} `json:"data"`
}

type RunnerResult struct {
	Data struct {
		Runner struct {
			ID          string `json:"id"`
			Description string `json:"description"`
			Status      string `json:"status"`
			RunnerType  string `json:"runnerType"`
			Managers    struct {
				Nodes []struct {
					SystemID  string `json:"systemId"`
					IPAddress string `json:"ipAddress"`
					Version   string `json:"version"`
					Revision  string `json:"revision"`
				} `json:"nodes"`
			} `json:"managers"`
		} `json:"runner"`
	}
}

func GraphQuery() {
	// Reference:
	// https://docs.gitlab.com/api/graphql/

	// TODO: How to get all runners?
	// https://docs.gitlab.com/api/graphql/reference/#queryrunners

	// Explorer, also exists on self-hosted:
	// https://gitlab.com/-/graphql-explorer

	c := gitlab.NewEnvironment()
	runner := &RunnerResult{}
	// Not sure how to get all runners, but this works, and IDs can be loaded the other API way
	c.Query(
		`query {
  runner(id: "gid://gitlab/Ci::Runner/1") {
    id description status runnerType
    managers { nodes { systemId ipAddress version revision } }
  }
}`,
		&runner,
	)
	fmt.Printf("Response: %+v\n", runner)
	fmt.Printf("Runner: %+v\n", c.GraphRunner(1))
}
