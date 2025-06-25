package collect

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/bubbletea/model/monitor/fetch"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/system"
	timeLibrary "github.com/funtimecoding/go-library/pkg/time"
	"time"
)

func Loop() {
	fmt.Printf("Commands: %s\n", join.Comma(fetch.List()))
	loopCheck(time.Now())
	ticker := time.NewTicker(5 * time.Minute)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				loopCheck(t)
			}
		}
	}()
	system.KillSignalBlock()
	done <- true
}

func loopCheck(t time.Time) {
	fmt.Printf(
		"Time: %s\n",
		t.Format(timeLibrary.DateMinute),
	)
	Check()
}
