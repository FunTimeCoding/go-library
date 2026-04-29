package gosentry

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors/sentry"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func showIssue(shortID string) {
	c := sentry.NewEnvironment()
	org := environment.Required(constant.OrganizationEnvironment)
	i := c.MustIssueByShortIdentifier(org, shortID)

	if i == nil {
		fmt.Printf("Issue not found: %s\n", shortID)

		return
	}

	r := i.Raw
	fmt.Printf("Issue:    %s\n", r.ShortID)
	fmt.Printf("Title:    %s\n", r.Title)
	fmt.Printf("Project:  %s\n", r.Project.Name)
	fmt.Printf("Link:     %s\n", r.Permalink)
	fmt.Printf("Status:   %s\n", r.Status)
	fmt.Printf("Level:    %s\n", r.Level)
	fmt.Printf("Events:   %s\n", r.Count)

	if r.FirstSeen != nil {
		fmt.Printf(
			"First:    %s\n",
			r.FirstSeen.Format("2006-01-02 15:04"),
		)
	}

	if r.LastSeen != nil {
		fmt.Printf(
			"Last:     %s\n",
			r.LastSeen.Format("2006-01-02 15:04"),
		)
	}

	if r.Culprit != "" {
		fmt.Printf("Culprit:  %s\n", r.Culprit)
	}

	e := c.MustLatestEvent(org, r.ID)
	fmt.Println()
	fmt.Printf("Latest Event: %s\n", e.EventID)

	if e.DateCreated != nil {
		fmt.Printf(
			"Date:     %s\n",
			e.DateCreated.Format("2006-01-02 15:04"),
		)
	}

	if len(e.Entries) > 0 {
		fmt.Println()
		printEventEntries(e.Entries)
	}
}
