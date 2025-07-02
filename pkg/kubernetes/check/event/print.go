package event

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/kubernetes/check/event/option"
	"github.com/funtimecoding/go-library/pkg/kubernetes/client"
	"github.com/funtimecoding/go-library/pkg/kubernetes/constant"
	"time"
)

func Print(o *option.Event) {
	k := client.NewEnvironment()
	cleanup(k, o, 7*24*time.Hour)

	if o.Notation {
		printNotation(k, o)

		return
	}

	f := constant.Dense.Copy()

	for _, e := range k.EventsSimple(false, true) {
		fmt.Println(e.Format(f))
	}
}
