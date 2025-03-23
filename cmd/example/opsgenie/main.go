package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/opsgenie"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
	"github.com/funtimecoding/go-library/pkg/separator"
	"github.com/funtimecoding/go-library/pkg/strings/split/key_value"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"strings"
)

func main() {
	pflag.Bool(argument.Create, false, "Create alert")
	pflag.String(argument.User, "", "User email for alert")
	pflag.String(argument.Text, "", "Alert name")
	pflag.String(argument.Close, "", "Alert ID")
	argument.ParseAndBind()
	c := opsgenie.NewEnvironment()
	c.TeamMap().AddKey("Infinite Loopsies", "INF")
	c.ShortAlert(
		func(s string) string {
			switch s {
			case constant.HighMemoryUsage:
				return "Memory"
			}

			return s
		},
	)
	c.ShortUser(
		func(s string) string {
			if strings.Contains(s, separator.At) {
				k, _ := key_value.At(s)

				return k
			}

			return s
		},
	)
	c.DescriptionToName(
		func(s string) string {
			return s
		},
	)
	c.TagsToTeam(
		func(s []string) string {
			return ""
		},
	)

	if viper.GetBool(argument.Create) {
		c.Create(
			viper.GetString(argument.User),
			viper.GetString(argument.Text),
		)

		return
	}

	if i := viper.GetString(argument.Close); i != "" {
		c.Close(i)

		return
	}

	f := option.ExtendedColor.Copy()
	open := c.Open()

	for _, a := range open {
		fmt.Println(a.Format(f))
	}

	if len(open) == 0 {
		fmt.Println("No open alerts")
	}
}
