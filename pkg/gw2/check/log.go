package check

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/gw2"
	"github.com/funtimecoding/go-library/pkg/gw2/check/aleeva_report"
	"github.com/funtimecoding/go-library/pkg/gw2/check/exceptions"
	"github.com/funtimecoding/go-library/pkg/gw2/check/guilds"
	"github.com/funtimecoding/go-library/pkg/gw2/check/meta"
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

func Log(
	path string,
	tag string,
) {
	m := meta.Parse("tmp/meta.json")
	var atRiskCutOff string
	var currentTeam string
	var metaToken string
	var matchUpStart time.Time

	if tag == "" && m.AllianceTag != "" {
		tag = m.AllianceTag
	}

	atRiskCutOff = m.AtRiskCutOff
	metaToken = m.Token
	currentTeam = m.Team
	matchUpStart = timeLibrary.Parse("2006-01-02", m.LinkStartDate)
	matchUpStart = time.Date(
		matchUpStart.Year(),
		matchUpStart.Month(),
		matchUpStart.Day(),
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
	token := environment.GetDefault(gw2.TokenEnvironment, metaToken)

	if token == "" {
		fmt.Println("No GW2 API token found")

		return
	}

	gw2.ImportAleevaFiles()
	c := gw2.New(token)
	members := gw2.MembersOfGuild(c, tag)
	fmt.Printf("Members count: %d\n", len(members))
	var onTeamAccounts []string
	var onDiscordAccounts []string
	var confirmedNotOnTeamAccounts []string
	files := system.FilesMatching(constant.Temporary, gw2.AleevaPrefix)
	var atRiskMembers []string

	for _, member := range members {
		atRiskMembers = append(atRiskMembers, member)

		if member == atRiskCutOff {
			break
		}
	}

	for _, file := range files {
		fmt.Printf("Aleeva file: %s\n", file)
	}

	latest := gw2.LatestAleevaFile(files)
	aleevaPath := system.Join(constant.Temporary, latest)
	fmt.Printf("Latest Aleeva file: %s\n", aleevaPath)

	t := tablewriter.NewWriter(os.Stdout)
	t.SetHeader([]string{"Name", "Account(s)", "Team(s)"})
	var verifiedAccounts []string
	var tableRowCount int

	for _, discordUser := range aleeva_report.Parse(aleevaPath) {
		if len(discordUser.WvwTeams) == 0 || len(discordUser.Gw2Accounts) == 0 {
			continue
		}

		if !gw2.IsAllianceMember(members, discordUser.Gw2Accounts) {
			continue
		}

		verifiedAccounts = append(
			verifiedAccounts,
			discordUser.Gw2Accounts...,
		)
		var teams []string

		for _, team := range discordUser.WvwTeams {
			teams = append(teams, team)
		}

		if slices.Contains(teams, currentTeam) {
			for account, teamId := range discordUser.WvwTeams {
				if teamId == currentTeam {
					onTeamAccounts = append(onTeamAccounts, account)
				}
			}

			continue
		}

		tableRowCount++
		t.Append(
			[]string{
				discordUser.DiscordName,
				join.Comma(discordUser.Gw2Accounts),
				join.Comma(teams),
			},
		)

		for account, teamId := range discordUser.WvwTeams {
			if teamId != currentTeam {
				confirmedNotOnTeamAccounts = append(
					confirmedNotOnTeamAccounts,
					account,
				)
			}
		}

		for account, teamId := range discordUser.WvwTeams {
			if teamId != currentTeam {
				onDiscordAccounts = append(onDiscordAccounts, account)
			}
		}
	}

	if tableRowCount > 0 {
		fmt.Println("Not on team members:")
		t.Render()
	}

	fmt.Printf("Members: %s\n", join.Comma(members))
	logs := log.NewSlice(gw2.ParseLogs(system.ReadBytes(path), false))
	var unverifiedAccounts []string

	for _, member := range members {
		if !slices.Contains(verifiedAccounts, member) {
			unverifiedAccounts = append(unverifiedAccounts, member)
		}
	}

	seenSlice := log.LastSeenPerMemberSlice(members, logs, &matchUpStart)
	fmt.Printf("Last seen count: %d\n", len(seenSlice))

	for _, v := range seenSlice {
		fmt.Printf(
			"Last seen: %s: %s\n",
			v.Name,
			v.Time.Format(timeLibrary.DateMinute),
		)
	}

	var exceptionNames []string

	for _, element := range exceptions.Parse("tmp/exception.json") {
		exceptionNames = append(exceptionNames, element.Name)
	}

	var neverSeen []string
	var foundExceptions []string
	seenMap := log.LastSeenPerMemberMap(members, logs, &matchUpStart)

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
	guildReport := guilds.Parse("tmp/guild.json")

	if false {
		// To maintain guilds.json
		for guild, guildMembers := range guildReport {
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
		labels := GuildsOfMember(guildReport, name)

		if len(labels) == 0 {
			labels = append(labels, "no-guild")
		}

		if slices.Contains(onTeamAccounts, name) {
			labels = append(labels, "on-current-team")
		}

		if slices.Contains(onDiscordAccounts, name) {
			labels = append(labels, "on-discord")
		}

		if slices.Contains(unverifiedAccounts, name) {
			labels = append(labels, "unverified")
		}

		if slices.Contains(atRiskMembers, name) {
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

	for _, element := range seenDays {
		seenMembers = append(seenMembers, element.Name)

		if len(element.Days) < 4 {
			lessThanFour = append(lessThanFour, element.Name)
		}
	}

	fmt.Printf("  Less than four days count: %d\n", len(lessThanFour))
	var neverSeenDays []string
	var exceptionCount int
	var noAleevaOrNotOnTeamCount int
	var confirmedNotOnTeamCount int

	for _, member := range members {
		if !slices.Contains(onTeamAccounts, member) {
			noAleevaOrNotOnTeamCount++
		}

		if slices.Contains(confirmedNotOnTeamAccounts, member) {
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

	for _, element := range seenDays {
		labels := GuildsOfMember(guildReport, element.Name)

		if len(labels) == 0 {
			labels = append(labels, "no-guild")
		}

		if slices.Contains(atRiskMembers, element.Name) {
			labels = append(labels, "at-risk")
		}

		fmt.Printf(
			"Seen days: %d %s %s\n",
			len(element.Days),
			element.Name,
			join.Space(labels...),
		)
	}

	for _, element := range neverSeenDays {
		labels := GuildsOfMember(guildReport, element)

		if len(labels) == 0 {
			labels = append(labels, "no-guild")
		}

		if slices.Contains(atRiskMembers, element) {
			labels = append(labels, "at-risk")
		}

		fmt.Printf("Seen days: 0 %s %s\n", element, join.Space(labels...))
	}
}
