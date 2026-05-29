package client_configuration

import (
	"k8s.io/client-go/tools/clientcmd"
	"sort"
)

func Contexts() ([]string, error) {
	c, e := clientcmd.LoadFromFile(path())

	if e != nil {
		return nil, e
	}

	var result []string

	for name := range c.Contexts {
		result = append(result, name)
	}

	sort.Strings(result)

	return result, nil
}
