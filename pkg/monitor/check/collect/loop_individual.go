package collect

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/bubbletea/model/monitor/fetch"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/system"
	timeLibrary "github.com/funtimecoding/go-library/pkg/time"
	"k8s.io/apimachinery/pkg/util/duration"
	"time"
)

func LoopIndividual() {
	fmt.Printf("Commands: %s\n", join.Comma(fetch.List()))
	byName := make(map[string]*LastRun)
	tick(byName, time.Now())
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				tick(byName, t)
			}
		}
	}()
	system.KillSignalBlock()
	done <- true
}

type LastRun struct {
	Command string
	Time    time.Time
}

func tick(
	byName map[string]*LastRun,
	t time.Time,
) {
	for _, s := range fetch.Settings {
		if r, okay := byName[s.Command]; okay {
			if t.Sub(r.Time) >= s.Interval {
				r.Time = t
				runTime(s, t)
			}
		} else {
			byName[s.Command] = &LastRun{Command: s.Command, Time: t}
			runTime(s, t)
		}
	}
}

func runTime(
	s *fetch.Setting,
	t time.Time,
) {
	fmt.Printf(
		"%s %s %s\n",
		t.Format(timeLibrary.DateSecond),
		s.Command,
		duration.HumanDuration(s.Interval),
	)
	fetch.Run(s.Command)
}
