package check

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gw2"
	"github.com/funtimecoding/go-library/pkg/gw2/check/aleeva_report"
	"github.com/funtimecoding/go-library/pkg/gw2/check/exceptions"
	"github.com/funtimecoding/go-library/pkg/gw2/check/guilds"
	"github.com/funtimecoding/go-library/pkg/gw2/constant"
	"github.com/funtimecoding/go-library/pkg/gw2/log_manager/log"
	"github.com/funtimecoding/go-library/pkg/strings/contains"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/system"
	systemConstant "github.com/funtimecoding/go-library/pkg/system/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	timeLibrary "github.com/funtimecoding/go-library/pkg/time"
	"github.com/olekukonko/tablewriter"
	"os"
	"slices"
	"sort"
	"time"
)

func Log(
	path string,
	tag string,
) {
	if tag == "" {
		tag = environment.Exit(constant.AllianceEnvironment)
	}

	atRiskCutOff := environment.Exit(constant.AtRiskCutOffEnvironment)
	currentTeam := environment.Exit(constant.TeamEnvironment)
	start := timeLibrary.Parse(
		"2006-01-02",
		environment.Exit(constant.LinkStartDateEnvironment),
	)
	matchUpStart := time.Date(
		start.Year(),
		start.Month(),
		start.Day(),
		20, // not 18, because log times are in UTC too
		0,
		0,
		0,
		time.UTC,
	)
	fmt.Printf(
		"Match-up start: %s\n",
		matchUpStart.Format(timeLibrary.DateMinute),
	)
	fmt.Printf("Alliance tag: %s\n", tag)
	fmt.Printf("Team: %s\n", currentTeam)
	fmt.Printf("At-risk cut-off member: %s\n", atRiskCutOff)

	gw2.ImportAleevaFiles()
	members := gw2.MembersOfGuild(gw2.NewEnvironment(), tag)
	fmt.Printf("Members count: %d\n", len(members))
	var onTeam []string
	var onDiscord []string
	var confirmedNotOnTeam []string
	var atRisk []string

	for _, member := range members {
		atRisk = append(atRisk, member)

		if member == atRiskCutOff {
			break
		}
	}

	aleevaPath := system.Join(
		systemConstant.Temporary,
		gw2.LatestAleevaFile(
			system.FilesMatching(
				systemConstant.Temporary,
				constant.AleevaPrefix,
			),
		),
	)
	fmt.Printf("Latest Aleeva file: %s\n", aleevaPath)

	var exceptionNames []string

	for _, e := range exceptions.Parse("tmp/exception.json") {
		exceptionNames = append(exceptionNames, e.Name)
	}

	t := tablewriter.NewWriter(os.Stdout)
	t.Header([]string{"Name", "Account(s)", "Team(s)"})
	var verifiedAccounts []string
	var rowCount int

	for _, r := range aleeva_report.Parse(aleevaPath) {
		if len(r.WvwTeams) == 0 || len(r.Gw2Accounts) == 0 {
			continue
		}

		if !gw2.IsAllianceMember(members, r.Gw2Accounts) {
			continue
		}

		if contains.Any(r.Gw2Accounts, exceptionNames) {
			continue
		}

		verifiedAccounts = append(verifiedAccounts, r.Gw2Accounts...)
		var teams []string

		for _, team := range r.WvwTeams {
			teams = append(teams, team)
		}

		if slices.Contains(teams, currentTeam) {
			for account, teamId := range r.WvwTeams {
				if teamId == currentTeam {
					onTeam = append(onTeam, account)
				}
			}

			continue
		}

		rowCount++
		errors.PanicOnError(
			t.Append(
				[]string{
					r.DiscordName,
					join.Comma(r.Gw2Accounts),
					join.Comma(teams),
				},
			),
		)

		for account, teamId := range r.WvwTeams {
			if teamId != currentTeam {
				confirmedNotOnTeam = append(confirmedNotOnTeam, account)
			}
		}

		for account, teamId := range r.WvwTeams {
			if teamId != currentTeam {
				onDiscord = append(onDiscord, account)
			}
		}
	}

	if rowCount > 0 {
		fmt.Printf("Not on team members (%d):\n", rowCount)
		errors.PanicOnError(t.Render())
	}

	fmt.Printf("Members: %s\n", join.Comma(members))
	logs := log.NewSlice(gw2.ParseLogs(system.ReadBytes(path), false))
	var unverified []string

	for _, member := range members {
		if !slices.Contains(verifiedAccounts, member) {
			unverified = append(unverified, member)
		}
	}

	seen := log.LastSeenPerMemberSlice(members, logs, &matchUpStart)
	fmt.Printf("Last seen since relink count: %d\n", len(seen))

	for _, v := range seen {
		fmt.Printf(
			"Last seen: %s: %s\n",
			v.Name,
			v.Time.Format(timeLibrary.DateMinute),
		)
	}

	var neverSeen []string
	var foundExceptions []string
	seenMap := log.LastSeenPerMemberMap(members, logs, &matchUpStart)

	for _, member := range members {
		if _, okay := seenMap[member]; !okay {
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

	for _, e := range exceptionNames {
		if !slices.Contains(members, e) {
			uselessException = append(uselessException, e)
			fmt.Printf("Useless exception: %s\n", e)
		}
	}

	fmt.Printf("Useless exceptions: %d\n", len(uselessException))
	guildReport := guilds.Parse("tmp/guild.json")

	if false {
		// To maintain guilds.json
		for guild, guildMembers := range guildReport {
			for _, e := range guildMembers {
				if !slices.Contains(members, e) {
					fmt.Printf(
						"Guild member not found: %s %s\n",
						guild,
						e,
					)
				}
			}
		}
	}

	sort.Strings(foundExceptions)
	fmt.Printf(
		"Exceptions (%d): %s\n",
		len(foundExceptions),
		join.Comma(foundExceptions),
	)

	fmt.Printf("Never seen count: %d\n", len(neverSeen))
	sort.Strings(neverSeen)

	for _, name := range neverSeen {
		labels := GuildsOfMember(guildReport, name)

		if len(labels) == 0 {
			labels = append(labels, "no-guild")
		}

		if slices.Contains(onTeam, name) {
			labels = append(labels, "on-current-team")
		}

		if slices.Contains(onDiscord, name) {
			labels = append(labels, "on-discord")
		}

		if slices.Contains(unverified, name) {
			labels = append(labels, "unverified")
		}

		if slices.Contains(atRisk, name) {
			labels = append(labels, "at-risk")
		}

		fmt.Printf("Never seen: %s %s\n", name, join.Space(labels...))
	}

	daysAgo := 56
	ago := time.Now().AddDate(0, 0, -daysAgo)
	seenDays := log.SeenDaysPerMemberSlice(members, logs, &ago)
	fmt.Printf("Total members: %d\n", len(members))
	fmt.Printf(
		"  Seen days count since %s (%d days): %d\n",
		ago.Format(timeLibrary.DateMinute),
		daysAgo,
		len(seenDays),
	)
	var seenMembers []string
	var lessThanFour []string

	for _, e := range seenDays {
		seenMembers = append(seenMembers, e.Name)

		if len(e.Days) < 4 {
			lessThanFour = append(lessThanFour, e.Name)
		}
	}

	fmt.Printf("  Less than four days count: %d\n", len(lessThanFour))
	var neverSeenDays []string
	var exceptionCount int
	var noAleevaOrNotOnTeamCount int
	var confirmedNotOnTeamCount int

	for _, member := range members {
		if !slices.Contains(onTeam, member) {
			noAleevaOrNotOnTeamCount++
		}

		if slices.Contains(confirmedNotOnTeam, member) {
			confirmedNotOnTeamCount++
		}

		if slices.Contains(exceptionNames, member) {
			exceptionCount++

			continue
		}

		if !slices.Contains(seenMembers, member) {
			neverSeenDays = append(neverSeenDays, member)
		}
	}

	fmt.Printf("  Never seen days count: %d\n", len(neverSeenDays))
	fmt.Printf("    Exception count: %d\n", exceptionCount)
	fmt.Printf(
		"    Confirmed not on team count: %d\n",
		confirmedNotOnTeamCount,
	)
	fmt.Printf(
		"    No Aleeva count: %d\n",
		noAleevaOrNotOnTeamCount-confirmedNotOnTeamCount,
	)

	for _, e := range seenDays {
		labels := GuildsOfMember(guildReport, e.Name)

		if len(labels) == 0 {
			labels = append(labels, "no-guild")
		}

		if slices.Contains(atRisk, e.Name) {
			labels = append(labels, "at-risk")
		}

		fmt.Printf(
			"Seen days: %d %s %s\n",
			len(e.Days),
			e.Name,
			join.Space(labels...),
		)
	}

	for _, e := range neverSeenDays {
		labels := GuildsOfMember(guildReport, e)

		if len(labels) == 0 {
			labels = append(labels, "no-guild")
		}

		if slices.Contains(atRisk, e) {
			labels = append(labels, "at-risk")
		}

		fmt.Printf("Seen days: 0 %s %s\n", e, join.Space(labels...))
	}
}
