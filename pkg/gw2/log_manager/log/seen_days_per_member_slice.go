package log

import (
	"slices"
	"sort"
	"time"
)

func SeenDaysPerMemberSlice(
	members []string,
	logs []*Log,
	since *time.Time,
) []*SeenDaysMember {
	seenMap := make(map[string]*SeenDaysMember)

	for _, member := range members {
		for _, log := range logs {
			if slices.Contains(log.Accounts, member) {
				if since != nil && log.Time.Before(*since) {
					continue
				}

				if _, ok := seenMap[member]; !ok {
					seenMap[member] = &SeenDaysMember{Name: member}
				}

				seenMap[member].Fights++

				if !DaysContains(seenMap[member].Days, log.Time) {
					seenMap[member].Days = append(
						seenMap[member].Days,
						log.Time,
					)
				}
			}
		}
	}

	var result []*SeenDaysMember

	for _, seen := range seenMap {
		result = append(result, seen)
	}

	sort.SliceStable(
		result,
		func(
			i int,
			j int,
		) bool {
			return len(result[i].Days) > len(result[j].Days)
		},
	)

	return result
}
