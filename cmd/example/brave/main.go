package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/brave"
	"github.com/funtimecoding/go-library/pkg/brave/constant"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/system"
	"path/filepath"
	"strings"
)

func main() {
	if true {
		profiles()
	}

	if false {
		for _, p := range brave.Profiles() {
			fmt.Printf("Profile: %s\n", p)

			if p == constant.Profile2 {
				brave.OpenProfile(p)
			}
		}
	}
}

type Preferences struct {
	Cookie struct {
		Accounts string `json:"last_list_accounts_data"`
	} `json:"gaia_cookie"`
}

func parsePreferences(profile string) *Preferences {
	var result Preferences
	notation.DecodeBytes(
		system.ReadBytes(
			filepath.Join(
				filepath.Join(system.Home(), constant.BraveSettings),
				profile,
				constant.PreferencesFile,
			),
		),
		&result,
	)

	return &result
}

func profiles() {
	for _, f := range system.ReadDirectory(
		filepath.Join(system.Home(), constant.BraveSettings),
	) {
		if !f.IsDir() {
			continue
		}

		if f.Name() == constant.DefaultProfile ||
			strings.HasPrefix(f.Name(), constant.ProfilePrefix) {
			fmt.Printf("Profile: %+v\n", f.Name())
			fmt.Printf(
				"Account: %+v\n",
				parseAccounts(
					parsePreferences(f.Name()).Cookie.Accounts,
				)[0].Email,
			)
		}
	}
}

type Account struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func parseAccounts(s string) []*Account {
	var a []any
	notation.DecodeStrict(s, &a, false)
	var result []*Account

	for _, a2 := range a[1].([]any) {
		a3 := a2.([]any)
		result = append(
			result,
			&Account{Name: a3[2].(string), Email: a3[3].(string)},
		)
	}

	return result
}
