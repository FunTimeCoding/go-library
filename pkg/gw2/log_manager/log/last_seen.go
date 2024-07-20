package log

import "time"

func lastSeen(
	logs []*Log,
	member string,
) *time.Time {
	var result *time.Time

	for _, l := range logs {
		for _, account := range l.Accounts {
			if account == member {
				if result == nil || l.Time.After(*result) {
					result = &l.Time
				}
			}
		}
	}

	return result
}
