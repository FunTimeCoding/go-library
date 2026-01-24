package mark

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/generative/model_context/example/mark/option"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Main() {
	pflag.Bool(argument.Local, false, "Local STDIO")
	argument.ParseBind()
	o := option.New()
	o.Local = viper.GetBool(argument.Local)
	Run(o)
}
