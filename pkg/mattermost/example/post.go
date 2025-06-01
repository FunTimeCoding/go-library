package example

import "github.com/funtimecoding/go-library/pkg/mattermost"

func Post() {
	m := mattermost.NewEnvironment()
	m.PostDefault("Hello friend")
}
