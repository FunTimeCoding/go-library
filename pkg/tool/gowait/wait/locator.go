package wait

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gowait/wait/option"
	"io"
	"net/http"
	"strings"
	"time"
)

func Locator(o *option.Wait) {
	expect := argument.Required(argument.Contains)
	x, cancel := context.WithTimeout(context.Background(), o.Timeout)
	defer cancel()
	c := &http.Client{Timeout: 10 * time.Second}
	attempt := 0

	for {
		attempt++

		if o.Verbose {
			fmt.Printf("%s %d\n", o.Locator, attempt)
		}

		r, getFail := http.NewRequestWithContext(
			x,
			http.MethodGet,
			o.Locator,
			nil,
		)
		errors.PanicOnError(getFail)
		result, doFail := c.Do(r)

		if doFail != nil {
			if o.Verbose {
				fmt.Printf("Do fail: %v\n", doFail)
			}
		} else {
			body, readFail := io.ReadAll(result.Body)
			errors.LogClose(result.Body)

			if readFail != nil {
				if o.Verbose {
					fmt.Printf("Read fail: %v\n", readFail)
				}
			} else if result.StatusCode == http.StatusOK {
				content := string(body)

				if o.Verbose {
					fmt.Printf("Response: %s\n", content)
				}

				if strings.Contains(content, expect) {
					fmt.Println("Found")

					return
				}
			} else if o.Verbose {
				fmt.Printf("Status: %d\n", result.StatusCode)
			}
		}

		select {
		case <-x.Done():
			panic("timeout")
		case <-time.After(interval):
			// Continue
		}
	}
}
