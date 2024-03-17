package system

import (
	"fmt"
	"net"
	"strings"
)

func FindHostnames(instance string) ([]string, error) {
	var result []string
	var host string

	if strings.ContainsRune(instance, ':') {
		splitHost, _, splitFail := net.SplitHostPort(instance)

		if splitFail != nil {
			return result, fmt.Errorf("split: %s", splitFail)
		}

		host = splitHost
	} else {
		host = instance
	}

	if net.ParseIP(host) != nil {
		names, lookUpFail := net.LookupAddr(host)

		if lookUpFail != nil {
			if strings.HasSuffix(
				lookUpFail.Error(),
				"nodename nor servname provided, or not known",
			) {
				return result, nil
			}

			return result, fmt.Errorf("lookup: %s", lookUpFail)
		}

		result = names

		for i := range result {
			result[i] = strings.TrimSuffix(result[i], ".")
		}
	} else {
		result = []string{host}
	}

	return result, nil
}
