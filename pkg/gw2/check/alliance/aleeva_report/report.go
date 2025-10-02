package aleeva_report

type Report struct {
	DiscordName         string            `json:"discordName"`
	DiscordTag          string            `json:"discordTag"`
	DiscordId           string            `json:"discordId"`
	DiscordJoinTime     string            `json:"discordJoinTime"`
	DiscordCreationTime string            `json:"discordCreationTime"`
	Permissions         []string          `json:"permissions"`
	Gw2Accounts         []string          `json:"gw2Accounts"`
	ValidKeys           map[string]bool   `json:"validKeys"`
	InServerGuild       any               `json:"inServerGuild"`
	InFollowerGuild     map[string]bool   `json:"inFollowerGuild"`
	WvwTeams            map[string]string `json:"wvwTeams"`
	WvwGuilds           map[string]string `json:"wvwGuilds"`
	AccountCreation     map[string]string `json:"accountCreation"`
	LastOnline          map[string]string `json:"lastOnline"`
	IsAdmin             bool              `json:"isAdmin"`
	IsBoosting          bool              `json:"isBoosting"`
	IsTimedOut          bool              `json:"isTimedOut"`
	IsPending           bool              `json:"isPending"`
}
