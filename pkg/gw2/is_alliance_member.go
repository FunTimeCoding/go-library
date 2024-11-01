package gw2

import "slices"

func IsAllianceMember(
	allMembers []string,
	accounts []string,
) bool {
	for _, account := range accounts {
		if slices.Contains(allMembers, account) {
			return true
		}
	}

	return false
}
