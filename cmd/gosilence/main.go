package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/console/format"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	pflag.StringP(
		argument.Set,
		"n",
		"",
		"Create or update silence with name",
	)
	argument.ParseAndBind()
	c := alertmanager.NewEnvironment()

	if n := viper.GetString(argument.Set); n != "" {
		fmt.Printf("Set: %s\n", c.SimpleSilence(n))
	}

	silences := c.Silences(true)
	var active int

	for _, a := range silences {
		if a.State != constant.ActiveState {
			continue
		}

		active++
		fmt.Println(a.Format(format.ExtendedColor))
	}

	if active == 0 {
		fmt.Printf("No active silences, %d in total\n", len(silences))
	}
}
