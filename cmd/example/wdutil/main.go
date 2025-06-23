package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/macos/constant"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/run"
	timeLibrary "github.com/funtimecoding/go-library/pkg/time"
	"regexp"
	"strings"
	"time"
)

const NotAvailable = "n/a"

func main() {
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)
	go func() {
		past := NotAvailable

		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				r := collect()

				if past != r.Sequence {
					fmt.Printf(
						"%s change %s\n",
						t.Format(timeLibrary.DateSecond),
						r.Sequence,
					)
				}

				past = r.Sequence
			}
		}
	}()
	system.KillSignalBlock()
	done <- true
}

type Result struct {
	Sequence string
}

func collect() *Result {
	r := run.New()
	r.Start("sudo", constant.Wdutil, constant.Info)

	if false {
		fmt.Printf("Status: %s\n", r.OutputString)
	}

	sequence := parseKey(r.OutputString, "Channel Sequence")

	if false {
		fmt.Printf("Sequence: %s\n", sequence)
	}

	return &Result{Sequence: sequence}
}

func parseKey(
	output string,
	key string,
) string {
	m := regexp.MustCompile(
		fmt.Sprintf(`(?m)^\s*%s\s*:\s*(.+)$`, regexp.QuoteMeta(key)),
	).FindStringSubmatch(output)

	if len(m) >= 2 {
		return strings.TrimSpace(m[1])
	}

	return ""
}
