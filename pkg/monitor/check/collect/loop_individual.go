package collect

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/bubbletea/model/monitor/fetch"
	"github.com/funtimecoding/go-library/pkg/monitor/collector"
	"github.com/funtimecoding/go-library/pkg/monitor/item/constant"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/system"
	library "github.com/funtimecoding/go-library/pkg/time"
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
	for _, s := range constant.Collectors {
		if r, okay := byName[s.Name]; okay {
			if t.Sub(r.Time) >= s.Interval {
				r.Time = t
				runTime(s, t)
			}
		} else {
			byName[s.Name] = &LastRun{Command: s.Name, Time: t}
			runTime(s, t)
		}
	}
}

func runTime(
	s *collector.Collector,
	t time.Time,
) {
	fmt.Printf(
		"%s %s %s\n",
		t.Format(library.DateSecond),
		s.Name,
		duration.HumanDuration(s.Interval),
	)
	fetch.Run(s.Name)
}
