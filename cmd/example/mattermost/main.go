package main

import "github.com/funtimecoding/go-library/pkg/mattermost"

func main() {
	m := mattermost.NewEnvironment()
	m.PostDefault("Hello friend")
}
