package log

import "time"

func LastSeenPerMemberMap(
	members []string,
	logs []*Log,
	since *time.Time,
) map[string]time.Time {
	result := make(map[string]time.Time)

	for _, member := range members {
		if t := lastSeen(logs, member); t != nil {
			if since != nil && t.Before(*since) {
				continue
			}

			result[member] = *t
		}
	}

	return result
}
