package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/chat/mattermost"
)

func Team() {
	m := mattermost.NewEnvironment()

	for _, t := range m.MustTeams(m.MustMe().Id) {
		fmt.Printf("Team: %s %s\n", t.Id, t.Name)
	}
}
