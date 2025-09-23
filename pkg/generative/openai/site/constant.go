package site

const (
	// NewSelector not unique, requires index
	NewSelector = `a[data-testid="create-new-chat-button"]`

	ProfileSelector = `[data-testid="accounts-profile-button"]`

	SettingsSelector = `[data-testid="settings-menu-item"]`

	PersonalizeSelector = `[data-testid="personalization-tab"]`

	MemoriesSelector = `[class="btn relative btn-secondary btn-small"]`

	// CloseMemoriesSelector not unique, requires index
	CloseMemoriesSelector = `div[role="dialog"] [data-testid="close-button"]`

	CloseSettingsSelector = `div[role="tablist"] [data-testid="close-button"]`

	PromptSelector = `#prompt-textarea`
)

var usefulAttributes = []string{"data-testid", "aria-label"}
