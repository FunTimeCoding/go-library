package log

import "sort"

func LastSeenPerMemberSlice(
	members []string,
	logs []*Log,
) []LastSeenMember {
	var result []LastSeenMember

	for _, member := range members {
		if t := lastSeen(logs, member); t != nil {
			result = append(
				result,
				LastSeenMember{Name: member, Time: *t},
			)
		}
	}

	sort.SliceStable(
		result,
		func(
			i int,
			j int,
		) bool {
			return result[i].Time.Before(result[j].Time)
		},
	)

	return result
}
