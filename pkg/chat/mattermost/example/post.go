package example

import "github.com/funtimecoding/go-library/pkg/tool/common"

func Post() {
	m := common.Mattermost()
	m.PostDefault("Hello friend")
}
