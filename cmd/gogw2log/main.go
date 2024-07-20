package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/gw2"
	"github.com/funtimecoding/go-library/pkg/gw2/log_manager/log"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	timeLibrary "github.com/funtimecoding/go-library/pkg/time"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"sort"
	"time"
)

func main() {
	path := fmt.Sprintf(
		"%s\\AppData\\Local\\ArcdpsLogManager",
		system.Home(),
	)

	if false {
		readGuilds(fmt.Sprintf("%s\\ApiDataCache.json", path))
	}

	if true {
		readLogs(fmt.Sprintf("%s\\LogDataCache.json", path))
	}
}

func readLogs(path string) {
	pflag.StringP(argument.Tag, "g", "", "Guild tag")
	argument.ParseAndBind()
	c := gw2.New(environment.Get(gw2.TokenEnvironment, 1))
	tag := viper.GetString(argument.Tag)
	members := gw2.MembersOfGuild(c, tag)
	fmt.Printf("Members: %s\n", join.Comma(members))

	logs := log.NewSlice(gw2.ParseLogs(system.ReadBytes(path), false))

	if false {
		for _, l := range logs {
			fmt.Printf("Log: %+v\n", l.Time.Format(timeLibrary.DateMinute))
			fmt.Printf("  Accounts: %+v\n", join.Comma(l.Accounts))
		}
	}

	seenSlice := LastSeenPerMemberSlice(members, logs)
	fmt.Printf("Last seen count: %d\n", len(seenSlice))

	for _, v := range seenSlice {
		fmt.Printf(
			"Last seen: %s: %s\n",
			v.Name,
			v.Time.Format(timeLibrary.DateMinute),
		)
	}

	var neverSeen []string
	seenMap := LastSeenPerMemberMap(members, logs)

	for _, member := range members {
		if _, ok := seenMap[member]; !ok {
			neverSeen = append(neverSeen, member)
		}
	}

	fmt.Printf("Never seen count: %d\n", len(neverSeen))
	sort.Strings(neverSeen)

	for _, name := range neverSeen {
		fmt.Printf("Never seen: %s\n", name)
	}
}

type LastSeenMember struct {
	Name string
	Time time.Time
}

func LastSeenPerMemberSlice(
	members []string,
	logs []*log.Log,
) []LastSeenMember {
	var result []LastSeenMember

	for _, member := range members {
		if t := LastSeen(logs, member); t != nil {
			result = append(
				result, LastSeenMember{
					Name: member,
					Time: *t,
				},
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

func LastSeenPerMemberMap(
	members []string,
	logs []*log.Log,
) map[string]time.Time {
	result := make(map[string]time.Time)

	for _, member := range members {
		if t := LastSeen(logs, member); t != nil {
			result[member] = *t
		}
	}

	return result
}

func LastSeen(
	logs []*log.Log,
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

func readGuilds(path string) {
	for _, guild := range gw2.ParseGuilds(system.ReadFile(path)) {
		fmt.Printf("Guild: %s\n", guild.Name)
	}
}
