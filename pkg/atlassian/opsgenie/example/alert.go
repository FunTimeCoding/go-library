package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/internal"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Alert() {
	pflag.Bool(argument.Create, false, "Create alert")
	pflag.String(argument.User, "", "User email for alert")
	pflag.String(argument.Text, "", "Alert name")
	pflag.String(argument.Close, "", "Alert ID")
	argument.ParseBind()
	c := internal.Opsgenie()

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
	alerts := c.Open()

	for _, a := range alerts {
		fmt.Println(a.Format(f))
	}

	if len(alerts) == 0 {
		fmt.Println("No relevant alerts")
	}
}
