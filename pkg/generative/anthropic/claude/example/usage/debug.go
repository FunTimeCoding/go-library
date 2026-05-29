package usage

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/anthropic/claude"
	"sort"
	"time"
)

func Debug() {
	c := claude.New()
	sessions := c.Sessions()
	var all []*timestamped

	for _, s := range sessions {
		for _, entry := range c.UsageEntries(s.Identifier) {
			t := parseTimestamp(entry.Timestamp)

			if t.IsZero() {
				continue
			}

			all = append(all, &timestamped{time: t, entry: entry})
		}
	}

	sort.Slice(
		all,
		func(i, j int) bool {
			return all[i].time.Before(all[j].time)
		},
	)
	fmt.Printf("Total entries across all sessions: %d\n\n", len(all))
	window := 5 * time.Hour
	fmt.Println("Gaps >= 1h in last 24h:")
	cutoff24 := now().Add(-24 * time.Hour)

	for i := 1; i < len(all); i++ {
		if all[i].time.Before(cutoff24) {
			continue
		}

		gap := all[i].time.Sub(all[i-1].time)

		if gap >= time.Hour {
			marker := ""

			if gap >= window {
				marker = " *** BLOCK BREAK ***"
			}

			fmt.Printf(
				"  %s → %s  gap=%s%s\n",
				all[i-1].time.Format("Jan 02 15:04"),
				all[i].time.Format("15:04"),
				gap.Truncate(time.Second),
				marker,
			)
		}
	}

	fmt.Println()
	block := findActiveBlock(all)

	if block == nil {
		fmt.Println("no active block")

		return
	}

	fmt.Printf(
		"Active block: %s - %s (%d entries)\n",
		block.start.Format("15:04"),
		block.end.Format("15:04"),
		len(block.entries),
	)
	var totalCost float64

	for _, ts := range block.entries {
		totalCost += entryCost(
			normalizeModel(ts.entry.Model),
			ts.entry,
		)
	}

	fmt.Printf(
		"Cost: $%.2f / $140.00  (%.0f%%)\n",
		totalCost,
		(totalCost/140)*100,
	)
}
