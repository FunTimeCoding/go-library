package alliance

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Check() {
	path := fmt.Sprintf(
		"%s\\AppData\\Local\\ArcdpsLogManager",
		system.Home(),
	)

	if false {
		Guild(path)
	}

	if true {
		pflag.String(argument.Tag, "", "Guild tag")
		argument.ParseBind()
		Log(path, viper.GetString(argument.Tag))
	}
}
