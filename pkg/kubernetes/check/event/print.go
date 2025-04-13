package event

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/kubernetes/check/event/option"
	"github.com/funtimecoding/go-library/pkg/kubernetes/client"
	"github.com/funtimecoding/go-library/pkg/kubernetes/constant"
	"time"
)

func Print(o *option.Event) {
	c := client.NewEnvironment()

	if o.Notation {
		printNotation(c, o)

		return
	}

	f := constant.Dense.Copy()
	age := 7 * 24 * time.Hour

	if o.Clean {
		for _, e := range c.EventsSimple(true, false) {
			if deleteOld(c, e, age) {
				continue
			}
		}
	}

	for _, e := range c.EventsSimple(false, true) {
		if o.Clean {
			if deleteIrrelevant(c, e) {
				continue
			}

			if deleteOld(c, e, age) {
				continue
			}
		}

		fmt.Println(e.Format(f))
	}
}
