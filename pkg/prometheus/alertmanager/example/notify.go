package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/macos"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
	"github.com/funtimecoding/go-library/pkg/sound"
	"github.com/funtimecoding/go-library/pkg/system"
	"time"
)

func Notify() {
	c := alertmanager.NewEnvironment()
	s := &State{}
	stop := make(chan struct{})
	go worker(stop, c, s)
	system.KillSignalBlock()
	close(stop)
	fmt.Println("Stop")
}

type State struct {
	Loaded bool
	Alerts []*alert.Alert
}

func worker(
	stop <-chan struct{},
	c *alertmanager.Client,
	s *State,
) {
	for {
		select {
		case <-stop:
			fmt.Println("Stopped")

			return
		default:
			now := c.Alerts()
			add, stay, remove := difference(s.Alerts, now)
			s.Alerts = now

			for _, a := range add {
				fmt.Printf("Add: %s\n", a.Name)
			}

			if false {
				for _, a := range stay {
					fmt.Printf("Stay: %s\n", a.Name)
				}
			}

			for _, a := range remove {
				fmt.Printf("Remove: %s\n", a.Name)
			}

			if !s.Loaded {
				s.Loaded = true
			} else {
				if len(add) > 0 {
					sound.Play(sound.SosumiPath, 1.0, false)
				}

				for _, a := range add {
					var summary string

					if a.Summary == constant.None {
						summary = "no summary"
					} else {
						summary = a.Summary
					}

					macos.Notify("Alert", a.Name, summary)
				}
			}

			time.Sleep(10 * time.Second)
		}
	}
}

func difference(
	past []*alert.Alert,
	now []*alert.Alert,
) ([]*alert.Alert, []*alert.Alert, []*alert.Alert) {
	var add []*alert.Alert
	var stay []*alert.Alert
	var remove []*alert.Alert

	for _, n := range now {
		var found bool

		for _, o := range past {
			if n.MonitorIdentifier == o.MonitorIdentifier {
				found = true

				break
			}
		}

		if !found {
			add = append(add, n)
		}
	}

	for _, o := range past {
		var found bool

		for _, n := range now {
			if n.MonitorIdentifier == o.MonitorIdentifier {
				found = true

				break
			}
		}

		if found {
			stay = append(stay, o)
		} else {
			remove = append(remove, o)
		}
	}

	return add, stay, remove
}
