package log

import "time"

func LastSeenPerMemberMap(
	members []string,
	logs []*Log,
) map[string]time.Time {
	result := make(map[string]time.Time)

	for _, member := range members {
		if t := lastSeen(logs, member); t != nil {
			result[member] = *t
		}
	}

	return result
}
