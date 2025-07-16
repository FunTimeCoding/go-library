package response

type Node struct {
	Parameters struct {
		ChannelId    any      `json:"channelId,omitempty"`
		Message      string   `json:"message,omitempty"`
		Attachments  []any    `json:"attachments,omitempty"`
		OtherOptions struct{} `json:"otherOptions,omitempty"`
		Owner        struct {
			Rl    bool   `json:"__rl"`
			Mode  string `json:"mode"`
			Value string `json:"value"`
		} `json:"owner,omitempty"`
		Repository struct {
			Rl    bool   `json:"__rl"`
			Mode  string `json:"mode"`
			Value string `json:"value"`
		} `json:"repository,omitempty"`
		Events  []string `json:"events,omitempty"`
		Options struct {
			SystemMessage string `json:"systemMessage,omitempty"`
		} `json:"options,omitempty"`
		Content          string   `json:"content,omitempty"`
		Height           int      `json:"height,omitempty"`
		Width            int      `json:"width,omitempty"`
		PromptType       string   `json:"promptType,omitempty"`
		Text             string   `json:"text,omitempty"`
		AdditionalFields struct{} `json:"additionalFields,omitempty"`
		Updates          []string `json:"updates,omitempty"`
		Resource         string   `json:"resource,omitempty"`
		Operation        string   `json:"operation,omitempty"`
		GuildId          struct {
			Rl    bool   `json:"__rl"`
			Mode  string `json:"mode"`
			Value string `json:"value"`
		} `json:"guildId,omitempty"`
		Authentication string `json:"authentication,omitempty"`
	} `json:"parameters"`
	Type        string         `json:"type"`
	TypeVersion float64        `json:"typeVersion"`
	Position    []int          `json:"position"`
	Id          string         `json:"id"`
	Name        string         `json:"name"`
	Credentials map[string]any `json:"credentials,omitempty"`
	WebhookId   string         `json:"webhookId,omitempty"`
}
