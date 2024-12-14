package check

import "slices"

func GuildsOfMember(
	report map[string][]string,
	account string,
) []string {
	var result []string

	for guildName, guildMembers := range report {
		if slices.Contains(guildMembers, account) {
			result = append(result, guildName)
		}
	}

	return result
}
