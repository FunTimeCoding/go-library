package check

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/gw2"
	"github.com/funtimecoding/go-library/pkg/gw2/log_manager/log"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	timeLibrary "github.com/funtimecoding/go-library/pkg/time"
	"sort"
)

func Log(
	path string,
	tag string,
) {
	c := gw2.New(environment.Get(gw2.TokenEnvironment, 1))
	members := gw2.MembersOfGuild(c, tag)
	fmt.Printf("Members: %s\n", join.Comma(members))

	logs := log.NewSlice(gw2.ParseLogs(system.ReadBytes(path), false))

	if false {
		for _, l := range logs {
			fmt.Printf("Log: %+v\n", l.Time.Format(timeLibrary.DateMinute))
			fmt.Printf("  Accounts: %+v\n", join.Comma(l.Accounts))
		}
	}

	seenSlice := log.LastSeenPerMemberSlice(members, logs)
	fmt.Printf("Last seen count: %d\n", len(seenSlice))

	for _, v := range seenSlice {
		fmt.Printf(
			"Last seen: %s: %s\n",
			v.Name,
			v.Time.Format(timeLibrary.DateMinute),
		)
	}

	var neverSeen []string
	seenMap := log.LastSeenPerMemberMap(members, logs)

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
