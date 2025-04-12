package client

import (
	"github.com/funtimecoding/go-library/pkg/kubernetes/client/client_configuration"
	"slices"
	"sort"
	"strings"
)

func (c *Client) ConfigurationContexts(filterExtended bool) []string {
	var result []string

	for context := range client_configuration.Interface().Contexts {
		result = append(result, context)
	}

	if filterExtended {
		var filtered []string

		for _, element := range result {
			if slices.Contains(filtered, element) {
				continue
			}

			var isExtendedName bool

			for _, inner := range result {
				if element == inner {
					continue
				}

				if strings.HasPrefix(element, inner) {
					isExtendedName = true
				}
			}

			if isExtendedName {
				continue
			}

			filtered = append(filtered, element)
		}

		result = filtered
	}

	sort.Strings(result)

	return result
}
