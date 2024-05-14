package status

import (
	"regexp"
	"strings"
)

func Parse(s string) *Result {
	result := &Result{}

	if m := regexp.MustCompile(
		`Loaded: (.+)`,
	).FindStringSubmatch(s); len(m) > 1 {
		result.Loaded = strings.TrimSpace(m[1])
	}

	if m := regexp.MustCompile(
		`Active: (.+) since (.+);`,
	).FindStringSubmatch(s); len(m) > 1 {
		result.Active = strings.TrimSpace(m[1])
	}

	if m := regexp.MustCompile(
		`Main PID: (\d+)`,
	).FindStringSubmatch(s); len(m) > 1 {
		result.MainProcess = strings.TrimSpace(m[1])
	}

	if m := regexp.MustCompile(
		`Memory: (.+)`,
	).FindStringSubmatch(s); len(m) > 1 {
		result.Memory = strings.TrimSpace(m[1])
	}

	if m := regexp.MustCompile(
		`CPU: (.+)`,
	).FindStringSubmatch(s); len(m) > 1 {
		result.Processor = strings.TrimSpace(m[1])
	}

	return result
}
