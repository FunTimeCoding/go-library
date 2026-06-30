package goc

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/strings/upper"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"testing"
)

func TestEnvironment(t *testing.T) {
	if false {
		// Only works for the running process
		environment.Set(ConfigurationEnvironment, upper.Alfa)
		environment.Set(ArgumentEnvironment, upper.Bravo)
	}

	if false {
		// Not sure if this works
		setEnvironmentEscape(ConfigurationEnvironment, upper.Alfa)
		setEnvironmentEscape(ArgumentEnvironment, upper.Bravo)
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
