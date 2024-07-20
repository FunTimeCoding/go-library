package log

import "sort"

func SortByTime(v []*Log) {
	sort.SliceStable(
		v,
		func(
			i int,
			j int,
		) bool {
			return v[i].Time.Before(v[j].Time)
		},
	)
}
