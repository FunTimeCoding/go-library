package post

import (
	"github.com/mattermost/mattermost/server/public/model"
	"sort"
)

func Sort(v []*model.Post) {
	sort.SliceStable(
		v,
		func(
			i int,
			j int,
		) bool {
			return v[i].CreateAt < v[j].CreateAt
		},
	)
}
