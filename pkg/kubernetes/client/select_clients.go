package client

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/kubernetes/filter"
)

func (c *Client) selectClients(f *filter.Filter) []*Client {
	var result []*Client

	if len(c.clients) == 0 {
		result = append(result, c)
	} else {
		for _, client := range c.clients {
			if c.Verbose {
				fmt.Printf("select client: %s\n", client.cluster)
				fmt.Printf("filter: %+v\n", f)
			}

			if f == nil || f.ContainsCluster(client.cluster) {
				result = append(result, client)
			}
		}
	}

	return result
}
