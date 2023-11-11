package git

import "github.com/funtimecoding/go-library/pkg/errors"

type Remote struct {
	Name    string
	Locator string
}

func Remotes(path string) []*Remote {
	remotes, e := Open(path).Remotes()
	errors.FatalOnError(e)
	var result []*Remote

	for _, element := range remotes {
		c := element.Config()
		result = append(result, &Remote{Name: c.Name, Locator: c.URLs[0]})
	}

	return result
}
