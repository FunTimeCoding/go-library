package example

import "github.com/funtimecoding/go-library/internal"

func Post() {
	m := internal.Mattermost()
	m.PostDefault("Hello friend")
}
