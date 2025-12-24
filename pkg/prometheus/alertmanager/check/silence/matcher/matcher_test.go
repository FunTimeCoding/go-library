package matcher

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/silence"
	"github.com/prometheus/alertmanager/api/v2/models"
	"testing"
	"time"
)

func TestMatchesWithExactMatcher(t *testing.T) {
	now := time.Now()
	start := now.Add(-1 * time.Hour)
	end := now.Add(1 * time.Hour)
	tests := []struct {
		name          string
		matcherName   string
		matcherValue  string
		alertLabels   models.LabelSet
		expectedMatch bool
		description   string
	}{
		{
			name:         "exact match with existing label",
			matcherName:  "job",
			matcherValue: "test",
			alertLabels: models.LabelSet{
				"job":       "test",
				"alertname": "HighCPU",
			},
			expectedMatch: true,
			description:   "Should match when label exists and value equals",
		},
		{
			name:         "exact match with different value",
			matcherName:  "job",
			matcherValue: "test",
			alertLabels: models.LabelSet{
				"job":       "production",
				"alertname": "HighCPU",
			},
			expectedMatch: false,
			description:   "Should not match when label exists but value differs",
		},
		{
			name:          "exact match with missing label",
			matcherName:   "job",
			matcherValue:  "test",
			alertLabels:   models.LabelSet{"alertname": "HighCPU"},
			expectedMatch: false,
			description:   "Should not match when label is missing (treated as empty string)",
		},
		{
			name:          "exact match empty string with missing label",
			matcherName:   "job",
			matcherValue:  "",
			alertLabels:   models.LabelSet{"alertname": "HighCPU"},
			expectedMatch: true,
			description:   "Should match when matcher is empty and label is missing",
		},
		{
			name:          "exact match empty string with existing empty label",
			matcherName:   "job",
			matcherValue:  "",
			alertLabels:   models.LabelSet{"job": "", "alertname": "HighCPU"},
			expectedMatch: true,
			description:   "Should match when both matcher and label are empty",
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				isEqual := true
				isRegex := false
				s := createSilence(
					&start,
					&end,
					tt.matcherName,
					tt.matcherValue,
					&isEqual,
					&isRegex,
				)
				a := createAlert(tt.alertLabels)
				matches := Matches(s, []*alert.Alert{a}, now)
				matched := len(matches) > 0

				if matched != tt.expectedMatch {
					t.Errorf(
						"%s: expected match=%v, got match=%v",
						tt.description,
						tt.expectedMatch,
						matched,
					)
				}
			},
		)
	}
}

func TestMatchesWithNotEqualMatcher(t *testing.T) {
	now := time.Now()
	start := now.Add(-1 * time.Hour)
	end := now.Add(1 * time.Hour)
	tests := []struct {
		name          string
		matcherName   string
		matcherValue  string
		alertLabels   models.LabelSet
		expectedMatch bool
		description   string
	}{
		{
			name:         "not equal with matching value",
			matcherName:  "job",
			matcherValue: "test",
			alertLabels: models.LabelSet{
				"job":       "test",
				"alertname": "HighCPU",
			},
			expectedMatch: false,
			description:   "Should not match when values are equal",
		},
		{
			name:         "not equal with different value",
			matcherName:  "job",
			matcherValue: "test",
			alertLabels: models.LabelSet{
				"job":       "production",
				"alertname": "HighCPU",
			},
			expectedMatch: true,
			description:   "Should match when values differ",
		},
		{
			name:          "not equal with missing label",
			matcherName:   "job",
			matcherValue:  "test",
			alertLabels:   models.LabelSet{"alertname": "HighCPU"},
			expectedMatch: true,
			description:   "Should match when label is missing (empty != test)",
		},
		{
			name:          "not equal empty string with missing label",
			matcherName:   "job",
			matcherValue:  "",
			alertLabels:   models.LabelSet{"alertname": "HighCPU"},
			expectedMatch: false,
			description:   "Should not match when both are empty",
		},
		{
			name:         "not equal empty string with non-empty label",
			matcherName:  "job",
			matcherValue: "",
			alertLabels: models.LabelSet{
				"job":       "test",
				"alertname": "HighCPU",
			},
			expectedMatch: true,
			description:   "Should match when label is non-empty and matcher is empty",
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				isEqual := false
				isRegex := false
				s := createSilence(
					&start,
					&end,
					tt.matcherName,
					tt.matcherValue,
					&isEqual,
					&isRegex,
				)
				a := createAlert(tt.alertLabels)
				matches := Matches(s, []*alert.Alert{a}, now)
				matched := len(matches) > 0

				if matched != tt.expectedMatch {
					t.Errorf(
						"%s: expected match=%v, got match=%v",
						tt.description,
						tt.expectedMatch,
						matched,
					)
				}
			},
		)
	}
}

func TestMatchesWithRegexMatcher(t *testing.T) {
	now := time.Now()
	start := now.Add(-1 * time.Hour)
	end := now.Add(1 * time.Hour)
	tests := []struct {
		name          string
		matcherName   string
		matcherValue  string
		alertLabels   models.LabelSet
		expectedMatch bool
		description   string
	}{
		{
			name:         "regex match with matching value",
			matcherName:  "job",
			matcherValue: "test.*",
			alertLabels: models.LabelSet{
				"job":       "test-job",
				"alertname": "HighCPU",
			},
			expectedMatch: true,
			description:   "Should match when regex matches",
		},
		{
			name:         "regex match with non-matching value",
			matcherName:  "job",
			matcherValue: "test.*",
			alertLabels: models.LabelSet{
				"job":       "production",
				"alertname": "HighCPU",
			},
			expectedMatch: false,
			description:   "Should not match when regex doesn't match",
		},
		{
			name:          "regex match with missing label",
			matcherName:   "job",
			matcherValue:  "test.*",
			alertLabels:   models.LabelSet{"alertname": "HighCPU"},
			expectedMatch: false,
			description:   "Should not match when label is missing (empty doesn't match test.*)",
		},
		{
			name:          "regex match empty string with missing label",
			matcherName:   "job",
			matcherValue:  "",
			alertLabels:   models.LabelSet{"alertname": "HighCPU"},
			expectedMatch: true,
			description:   "Should match when regex is empty and label is missing",
		},
		{
			name:          "regex match wildcard with missing label",
			matcherName:   "job",
			matcherValue:  ".*",
			alertLabels:   models.LabelSet{"alertname": "HighCPU"},
			expectedMatch: true,
			description:   "Should match when regex is .* (matches empty string)",
		},
		{
			name:         "regex match wildcard with existing label",
			matcherName:  "job",
			matcherValue: ".*",
			alertLabels: models.LabelSet{
				"job":       "test",
				"alertname": "HighCPU",
			},
			expectedMatch: true,
			description:   "Should match when regex is .* (matches any string)",
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				isEqual := true
				isRegex := true
				s := createSilence(
					&start,
					&end,
					tt.matcherName,
					tt.matcherValue,
					&isEqual,
					&isRegex,
				)
				a := createAlert(tt.alertLabels)
				matches := Matches(s, []*alert.Alert{a}, now)
				matched := len(matches) > 0

				if matched != tt.expectedMatch {
					t.Errorf(
						"%s: expected match=%v, got match=%v",
						tt.description,
						tt.expectedMatch,
						matched,
					)
				}
			},
		)
	}
}

func TestMatchesWithNotRegexMatcher(t *testing.T) {
	now := time.Now()
	start := now.Add(-1 * time.Hour)
	end := now.Add(1 * time.Hour)
	tests := []struct {
		name          string
		matcherName   string
		matcherValue  string
		alertLabels   models.LabelSet
		expectedMatch bool
		description   string
	}{
		{
			name:         "not regex with matching value",
			matcherName:  "job",
			matcherValue: "test.*",
			alertLabels: models.LabelSet{
				"job":       "test-job",
				"alertname": "HighCPU",
			},
			expectedMatch: false,
			description:   "Should not match when regex matches",
		},
		{
			name:         "not regex with non-matching value",
			matcherName:  "job",
			matcherValue: "test.*",
			alertLabels: models.LabelSet{
				"job":       "production",
				"alertname": "HighCPU",
			},
			expectedMatch: true,
			description:   "Should match when regex doesn't match",
		},
		{
			name:          "not regex with missing label",
			matcherName:   "job",
			matcherValue:  "test.*",
			alertLabels:   models.LabelSet{"alertname": "HighCPU"},
			expectedMatch: true,
			description:   "Should match when label is missing (empty doesn't match test.*)",
		},
		{
			name:          "not regex wildcard with missing label",
			matcherName:   "job",
			matcherValue:  ".*",
			alertLabels:   models.LabelSet{"alertname": "HighCPU"},
			expectedMatch: false,
			description:   "Should not match when regex is .* (it matches empty string)",
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				isEqual := false
				isRegex := true
				s := createSilence(
					&start,
					&end,
					tt.matcherName,
					tt.matcherValue,
					&isEqual,
					&isRegex,
				)
				a := createAlert(tt.alertLabels)
				matches := Matches(s, []*alert.Alert{a}, now)
				matched := len(matches) > 0

				if matched != tt.expectedMatch {
					t.Errorf(
						"%s: expected match=%v, got match=%v",
						tt.description,
						tt.expectedMatch,
						matched,
					)
				}
			},
		)
	}
}

func TestMatchesWithMultipleMatchers(t *testing.T) {
	now := time.Now()
	start := now.Add(-1 * time.Hour)
	end := now.Add(1 * time.Hour)
	tests := []struct {
		name          string
		matchers      []*models.Matcher
		alertLabels   models.LabelSet
		expectedMatch bool
		description   string
	}{
		{
			name: "all matchers match",
			matchers: []*models.Matcher{
				createMatcher("alertname", "HighCPU", true, false),
				createMatcher("severity", "critical", true, false),
			},
			alertLabels: models.LabelSet{
				"alertname": "HighCPU",
				"severity":  "critical",
			},
			expectedMatch: true,
			description:   "Should match when all matchers match",
		},
		{
			name: "one matcher doesn't match",
			matchers: []*models.Matcher{
				createMatcher("alertname", "HighCPU", true, false),
				createMatcher("severity", "warning", true, false),
			},
			alertLabels: models.LabelSet{
				"alertname": "HighCPU",
				"severity":  "critical",
			},
			expectedMatch: false,
			description:   "Should not match when any matcher doesn't match",
		},
		{
			name: "combined positive and negative matchers",
			matchers: []*models.Matcher{
				createMatcher("alertname", "HighCPU", true, false),
				createMatcher("job", "test", false, false),
			},
			alertLabels: models.LabelSet{
				"alertname": "HighCPU",
				"job":       "production",
			},
			expectedMatch: true,
			description:   "Should match when positive and negative matchers both match",
		},
		{
			name: "negative matcher with missing label",
			matchers: []*models.Matcher{
				createMatcher("alertname", "HighCPU", true, false),
				createMatcher("job", "test", false, false),
			},
			alertLabels:   models.LabelSet{"alertname": "HighCPU"},
			expectedMatch: true,
			description:   "Should match when negative matcher matches missing label",
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				s := createSilenceWithMatchers(&start, &end, tt.matchers)
				a := createAlert(tt.alertLabels)
				matches := Matches(s, []*alert.Alert{a}, now)
				matched := len(matches) > 0

				if matched != tt.expectedMatch {
					t.Errorf(
						"%s: expected match=%v, got match=%v",
						tt.description,
						tt.expectedMatch,
						matched,
					)
				}
			},
		)
	}
}

func TestMatchesWithTimeWindow(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name          string
		start         time.Time
		end           time.Time
		expectedMatch bool
		description   string
	}{
		{
			name:          "silence is active",
			start:         now.Add(-1 * time.Hour),
			end:           now.Add(1 * time.Hour),
			expectedMatch: true,
			description:   "Should match when current time is within silence window",
		},
		{
			name:          "silence not started yet",
			start:         now.Add(1 * time.Hour),
			end:           now.Add(2 * time.Hour),
			expectedMatch: false,
			description:   "Should not match when silence hasn't started",
		},
		{
			name:          "silence already expired",
			start:         now.Add(-2 * time.Hour),
			end:           now.Add(-1 * time.Hour),
			expectedMatch: false,
			description:   "Should not match when silence has expired",
		},
		{
			name:          "silence starts exactly now",
			start:         now,
			end:           now.Add(1 * time.Hour),
			expectedMatch: true,
			description:   "Should match when silence starts exactly at current time",
		},
		{
			name:          "silence ends exactly now",
			start:         now.Add(-1 * time.Hour),
			end:           now,
			expectedMatch: false,
			description:   "Should not match when silence ends exactly at current time",
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				isEqual := true
				isRegex := false
				s := createSilence(
					&tt.start,
					&tt.end,
					"alertname",
					"HighCPU",
					&isEqual,
					&isRegex,
				)
				a := createAlert(models.LabelSet{"alertname": "HighCPU"})

				matches := Matches(s, []*alert.Alert{a}, now)
				matched := len(matches) > 0

				if matched != tt.expectedMatch {
					t.Errorf(
						"%s: expected match=%v, got match=%v",
						tt.description,
						tt.expectedMatch,
						matched,
					)
				}
			},
		)
	}
}

func TestMatchesWithMultipleAlerts(t *testing.T) {
	now := time.Now()
	start := now.Add(-1 * time.Hour)
	end := now.Add(1 * time.Hour)
	equal := true
	regex := false
	s := createSilence(
		&start,
		&end,
		"severity",
		"critical",
		&equal,
		&regex,
	)
	alerts := []*alert.Alert{
		createAlert(
			models.LabelSet{"alertname": "HighCPU", "severity": "critical"},
		),
		createAlert(
			models.LabelSet{
				"alertname": "HighMemory",
				"severity":  "warning",
			},
		),
		createAlert(
			models.LabelSet{"alertname": "DiskFull", "severity": "critical"},
		),
		createAlert(
			models.LabelSet{"alertname": "NetworkIssue", "severity": "info"},
		),
	}
	m := Matches(s, alerts, now)

	if len(m) != 2 {
		t.Errorf("Expected 2 matching alerts, got %d", len(m))
	}

	for _, a := range m {
		if a.Labels["severity"] != "critical" {
			t.Errorf(
				"Expected all matched alerts to have severity=critical, got %s",
				a.Labels["severity"],
			)
		}
	}
}

func createMatcher(
	name string,
	value string,
	equal bool,
	regex bool,
) *models.Matcher {
	return &models.Matcher{
		Name:    &name,
		Value:   &value,
		IsEqual: &equal,
		IsRegex: &regex,
	}
}

func createSilence(
	start *time.Time,
	end *time.Time,
	matcherName string,
	matcherValue string,
	equal *bool,
	regex *bool,
) *silence.Silence {
	matcher := createMatcher(matcherName, matcherValue, *equal, *regex)
	return createSilenceWithMatchers(start, end, []*models.Matcher{matcher})
}

func createSilenceWithMatchers(
	start *time.Time,
	end *time.Time,
	m []*models.Matcher,
) *silence.Silence {
	return &silence.Silence{
		Start: start,
		End:   end,
		Raw:   &models.GettableSilence{Silence: models.Silence{Matchers: m}},
	}
}

func createAlert(l models.LabelSet) *alert.Alert {
	return &alert.Alert{
		Labels: l,
		Raw:    &models.GettableAlert{Alert: models.Alert{Labels: l}},
	}
}
