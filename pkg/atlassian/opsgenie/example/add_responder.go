package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/internal"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func AddResponder() {
	argument.ParseBind()
	r := argument.RequiredPositional(0, "RESPONDER_NAME")
	c := internal.Opsgenie()
	f := option.ExtendedColor.Copy()

	for _, a := range c.Open() {
		fmt.Println(a.Format(f))

		if false {
			c.AddResponderUser(a, r)
		}

		if true {
			break
		}
	}
}
