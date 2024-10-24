package keepass

import "github.com/tobischo/gokeepasslib/v3"

func ByTitleRecursive(
	g []gokeepasslib.Group,
	title string,
) *gokeepasslib.Entry {
	for _, group := range g {
		for _, entry := range group.Entries {
			if entry.GetTitle() == title {
				return &entry
			}
		}

		if e := ByTitleRecursive(group.Groups, title); e != nil {
			return e
		}
	}

	return nil
}
