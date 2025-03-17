package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/maps"
	"github.com/funtimecoding/go-library/pkg/markup"
	alertmanagerConstant "github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
	"github.com/funtimecoding/go-library/pkg/strings/split/key_value"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"github.com/funtimecoding/go-library/pkg/system/run"
	"slices"
	"strings"
)

func main() {
	// TODO: alertmanager configuration selection tool
	//  1) list which configurations are available
	//  2) show which configuration is selected
	//  3) switch between them with ease

	// go install github.com/prometheus/alertmanager/cmd/amtool@latest

	argument.ParseAndBind()
	base := system.Join(
		system.Home(),
		constant.ConfigurationPath,
		alertmanagerConstant.AmtoolPath,
	)
	tool := system.Join(
		base,
		alertmanagerConstant.AmtoolConfiguration,
	)

	if !run.CommandExists(alertmanagerConstant.AmtoolCommand) {
		fmt.Println(
			"amtool missing: go install github.com/prometheus/alertmanager/cmd/amtool@latest",
		)

		return
	}

	if !system.FileExists(tool) {
		fmt.Printf("Missing: %s\n", tool)

		return
	}

	active := ReadConfiguration(tool)
	files := system.Files(base)
	var contexts []string
	contextLocator := make(map[string]string)

	for _, f := range files {
		if false {
			fmt.Printf("File: %s\n", f)
		}

		name, _ := key_value.Dot(f)

		if strings.HasPrefix(
			name,
			alertmanagerConstant.AmtoolConfigurationPrefix,
		) {
			context := strings.TrimPrefix(
				name,
				alertmanagerConstant.AmtoolConfigurationPrefix,
			)

			if !slices.Contains(contexts, context) {
				contexts = append(contexts, context)
				p := system.Join(base, f)
				c := ReadConfiguration(p)
				contextLocator[context] = c.Locator
			}
		}
	}

	if len(contexts) == 0 {
		fmt.Println("No contexts found")

		return
	}

	name := argument.Positional(0)

	if name == "" {
		for _, k := range maps.StringKeys(contextLocator) {
			v := contextLocator[k]

			if v == active.Locator {
				fmt.Printf("* %s\n", k)
			} else {
				fmt.Printf("  %s\n", k)
			}
		}

		return
	}

	if !slices.Contains(contexts, name) {
		fmt.Printf("Context not found: %s\n", name)

		return
	}

	if active.Locator == contextLocator[name] {
		fmt.Printf("Already active: %s\n", name)

		return
	}

	system.CopyFile(
		system.Join(
			base,
			fmt.Sprintf(
				"%s%s.yml",
				alertmanagerConstant.AmtoolConfigurationPrefix,
				name,
			),
		),
		tool,
	)
	fmt.Printf("Now active: %s\n", name)
}

type Configuration struct {
	Locator          string `yaml:"alertmanager.url"`
	WebConfiguration string `yaml:"http.config.file"`
}

func ReadConfiguration(path string) *Configuration {
	c := &Configuration{}
	content := system.ReadFile(path)
	markup.DecodeStrict(content, &c)

	return c
}
