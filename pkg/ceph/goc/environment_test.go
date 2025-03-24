package goc

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"testing"
)

func TestEnvironment(t *testing.T) {
	if false {
		// Only works for the running process
		environment.Set(ConfigurationEnvironment, strings.Alfa)
		environment.Set(ArgumentEnvironment, strings.Bravo)
	}

	if false {
		// Not sure if this works
		setEnvironmentEscape(ConfigurationEnvironment, strings.Alfa)
		setEnvironmentEscape(ArgumentEnvironment, strings.Bravo)
	}

	if false {
		// Not working
		fmt.Printf(
			"Get escape: %s\n",
			getEnvironmentEscape(ConfigurationEnvironment),
		)
		fmt.Printf(
			"Get escape: %s\n",
			getEnvironmentEscape(ArgumentEnvironment),
		)
	}
}
