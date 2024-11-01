package check

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/gw2"
	"github.com/funtimecoding/go-library/pkg/gw2/check/aleeva_report"
	"github.com/funtimecoding/go-library/pkg/gw2/check/exceptions"
	"github.com/funtimecoding/go-library/pkg/gw2/check/guilds"
	"github.com/funtimecoding/go-library/pkg/gw2/log_manager/log"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	timeLibrary "github.com/funtimecoding/go-library/pkg/time"
	"github.com/olekukonko/tablewriter"
	"os"
	"slices"
	"sort"
	"time"
)

var MatchUpStart = time.Date(
	2024,
	10,
	25,
	20, // not 18, because log times are in UTC too
	0,
	0,
	0,
	time.UTC,
)

const CurrentTeam = "Bloodstone Gulch"

func Log(
	path string,
	tag string,
) {
	gw2.ImportAleevaFiles()
	c := gw2.New(environment.Get(gw2.TokenEnvironment, 1))
	members := gw2.MembersOfGuild(c, tag)
	fmt.Printf("Members count: %d\n", len(members))
	var onTeamMembers []string
	var onDiscord []string
	files := system.FilesMatching(constant.Temporary, gw2.AleevaPrefix)

	for _, file := range files {
		fmt.Printf("Aleeva file: %s\n", file)
	}

	latest := gw2.LatestAleevaFile(files)
	aleevaPath := system.Join(constant.Temporary, latest)
	fmt.Printf("Latest Aleeva file: %s\n", aleevaPath)

	t := tablewriter.NewWriter(os.Stdout)
	t.SetHeader([]string{"Name", "Account(s)", "Team(s)"})
	var verifiedAccounts []string

	for _, discordUser := range aleeva_report.Parse(aleevaPath) {
		if len(discordUser.WvwTeams) == 0 || len(discordUser.Gw2Accounts) == 0 {
			continue
		}

		if !gw2.IsAllianceMember(members, discordUser.Gw2Accounts) {
			continue
		}

		verifiedAccounts = append(verifiedAccounts, discordUser.Gw2Accounts...)
		var teams []string

		for _, team := range discordUser.WvwTeams {
			teams = append(teams, team)
		}

		if slices.Contains(teams, CurrentTeam) {
			continue
		}

		t.Append(
			[]string{
				discordUser.DiscordName,
				join.Comma(discordUser.Gw2Accounts),
				join.Comma(teams),
			},
		)

		for account, teamId := range discordUser.WvwTeams {
			if teamId == CurrentTeam {
				onTeamMembers = append(onTeamMembers, account)
			} else {
				onDiscord = append(onDiscord, account)
			}
		}
	}

	fmt.Println("Not on team members:")
	t.Render()

	fmt.Printf("Members: %s\n", join.Comma(members))
	logs := log.NewSlice(gw2.ParseLogs(system.ReadBytes(path), false))
	var unverifiedAccounts []string

	for _, member := range members {
		if !slices.Contains(verifiedAccounts, member) {
			unverifiedAccounts = append(unverifiedAccounts, member)
		}
	}

	seenSlice := log.LastSeenPerMemberSlice(members, logs, &MatchUpStart)
	fmt.Printf("Last seen count: %d\n", len(seenSlice))

	for _, v := range seenSlice {
		fmt.Printf(
			"Last seen: %s: %s\n",
			v.Name,
			v.Time.Format(timeLibrary.DateMinute),
		)
	}

	var exceptionNames []string

	for _, element := range exceptions.Parse("tmp/exceptions.json") {
		exceptionNames = append(exceptionNames, element.Name)
	}

	var neverSeen []string
	var foundExceptions []string
	seenMap := log.LastSeenPerMemberMap(members, logs, &MatchUpStart)

	for _, member := range members {
		if _, ok := seenMap[member]; !ok {
			if slices.Contains(exceptionNames, member) {
				foundExceptions = append(foundExceptions, member)
			} else {
				neverSeen = append(neverSeen, member)
			}
		} else {
			if slices.Contains(exceptionNames, member) {
				fmt.Printf("Found anyway exception: %s\n", member)
			}
		}
	}

	var uselessException []string

	for _, element := range exceptionNames {
		if !slices.Contains(members, element) {
			uselessException = append(uselessException, element)
			fmt.Printf("Useless exception: %s\n", element)
		}
	}

	fmt.Printf("Useless exceptions: %d\n", len(uselessException))
	guildsReport := guilds.Parse("tmp/guilds.json")

	if false {
		// To maintain guilds.json
		for guild, guildMembers := range guildsReport {
			for _, element := range guildMembers {
				if !slices.Contains(members, element) {
					fmt.Printf(
						"Guild member not found: %s %s\n",
						guild,
						element,
					)
				}
			}
		}
	}

	fmt.Printf("Exceptions count: %d\n", len(foundExceptions))
	sort.Strings(foundExceptions)

	for _, name := range foundExceptions {
		fmt.Printf("Exception: %s\n", name)
	}

	fmt.Printf("Never seen count: %d\n", len(neverSeen))
	sort.Strings(neverSeen)

	for _, name := range neverSeen {
		var extra string

		for guildName, guildMembers := range guildsReport {
			if slices.Contains(guildMembers, name) {
				extra = " " + guildName
			}

			break
		}

		if slices.Contains(onTeamMembers, name) {
			extra = " on-current-team"
		}

		if slices.Contains(onDiscord, name) {
			extra = " on-discord"
		}

		if slices.Contains(unverifiedAccounts, name) {
			extra = " unverified"
		}

		if extra == "" {
			extra = " no-guild"
		}

		fmt.Printf("Never seen: %s%s\n", name, extra)
	}
}
