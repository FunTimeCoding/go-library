package dictionary

import (
	"bufio"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system"
	"strings"
)

func Read(file string) []Category {
	f := system.Open(file)
	defer errors.LogClose(f)
	var categories []Category
	var category *Category
	s := bufio.NewScanner(f)

	for s.Scan() {
		line := strings.TrimSpace(s.Text())

		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "# ") {
			if category != nil {
				categories = append(categories, *category)
			}

			category = &Category{Name: strings.TrimPrefix(line, "# ")}
		} else if category != nil {
			category.Words = append(category.Words, line)
		}
	}

	errors.PanicOnError(s.Err())

	if category != nil {
		categories = append(categories, *category)
	}

	return categories
}
