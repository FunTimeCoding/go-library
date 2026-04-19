package parse_elite

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/raid/elite"
	"github.com/funtimecoding/go-library/pkg/raid/elite_parser"
	"os"
	"path/filepath"
	"strings"
)

func PlayerStats() {
	argument.ParseBind()
	directory := argument.RequiredPositional(0, "DIRECTORY")
	entries, e := os.ReadDir(directory)
	errors.PanicOnError(e)
	var allStats []elite_parser.PlayerFightStat
	fightCount := 0
	skipped := 0

	for _, entry := range entries {
		if !strings.HasSuffix(entry.Name(), "_detailed_wvw_kill.json") {
			continue
		}

		path := filepath.Join(directory, entry.Name())
		data, readError := os.ReadFile(path)
		errors.PanicOnError(readError)
		var fight elite.Fight
		errors.PanicOnError(json.Unmarshal(data, &fight))

		if !elite_parser.IsValidFight(&fight) {
			skipped++

			continue
		}

		fightCount++
		allStats = append(allStats, elite_parser.Extract(&fight)...)
	}

	players := elite_parser.Aggregate(allStats)
	fmt.Printf(
		"Aggregated %d fights (%d skipped), %d player-profession entries\n\n",
		fightCount,
		skipped,
		len(players),
	)
	sorted := sortByDamage(players)
	printOffensiveTable(sorted)
	fmt.Println()
	printDefensiveTable(sorted)
	fmt.Println()
	printSupportTable(sorted)
	fmt.Println()
	printBoonTable(sorted)
}
