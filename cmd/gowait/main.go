package main

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"strings"
	"time"
)

func main() {
	pflag.String(argument.Contains, "", "String to wait for")
	pflag.Duration(argument.Timeout, 3*time.Minute, "")
	monitor.VerboseArgument()
	argument.ParseBind()
	l := argument.RequiredPositional(0, "LOCATOR")
	expect := argument.Required(argument.Contains)
	timeout := viper.GetDuration(argument.Timeout)
	verbose := viper.GetBool(argument.Verbose)
	x, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	c := &http.Client{Timeout: 10 * time.Second}
	attempt := 0

	for {
		attempt++

		if verbose {
			fmt.Printf("%s %d\n", l, attempt)
		}

		r, getFail := http.NewRequestWithContext(
			x,
			http.MethodGet,
			l,
			nil,
		)
		errors.PanicOnError(getFail)
		result, doFail := c.Do(r)

		if doFail != nil {
			if verbose {
				fmt.Printf("Do fail: %v\n", doFail)
			}
		} else {
			body, readFail := io.ReadAll(result.Body)
			errors.LogClose(result.Body)

			if readFail != nil {
				if verbose {
					fmt.Printf("Read fail: %v\n", readFail)
				}
			} else if result.StatusCode == http.StatusOK {
				content := string(body)

				if verbose {
					fmt.Printf("Response: %s\n", content)
				}

				if strings.Contains(content, expect) {
					fmt.Println("Found")

					return
				}
			} else if verbose {
				fmt.Printf("Status: %d\n", result.StatusCode)
			}
		}

		select {
		case <-x.Done():
			panic("timeout")
		case <-time.After(10 * time.Second):
			// Continue
		}
	}
}
