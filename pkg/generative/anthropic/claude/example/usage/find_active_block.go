package usage

import "time"

func findActiveBlock(all []*timestamped) *block {
	if len(all) == 0 {
		return nil
	}

	window := 5 * time.Hour
	n := now()
	var current *block

	for _, ts := range all {
		needNew := current == nil

		if !needNew {
			pastEnd := ts.time.Equal(current.end) ||
				ts.time.After(current.end)
			gap := len(current.entries) > 0 &&
				ts.time.Sub(
					current.entries[len(current.entries)-1].time,
				) >= window
			needNew = pastEnd || gap
		}

		if needNew {
			start := ts.time.Truncate(time.Hour)
			current = &block{
				start: start,
				end:   start.Add(window),
			}
		}

		current.entries = append(current.entries, ts)
	}

	if current != nil && current.end.After(n) {
		return current
	}

	return nil
}
