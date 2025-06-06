package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/gw2/check"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	path := fmt.Sprintf(
		"%s\\AppData\\Local\\ArcdpsLogManager",
		system.Home(),
	)

	if false {
		check.Guild(fmt.Sprintf("%s\\ApiDataCache.json", path))
	}

	if true {
		pflag.String(argument.Tag, "", "Guild tag")
		argument.ParseBind()
		check.Log(
			fmt.Sprintf("%s\\LogDataCache.json", path),
			viper.GetString(argument.Tag),
		)
	}
}
