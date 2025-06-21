package site

const (
	// NewSelector not unique, requires index
	NewSelector = `a[data-testid="create-new-chat-button"]`

	ProfileSelector = `[data-testid="profile-button"]`

	SettingsSelector = `[data-testid="settings-menu-item"]`

	PersonalizeSelector = `[data-testid="personalization-tab"]`

	MemoriesSelector = `[class="flex items-center justify-center"]`

	// CloseMemoriesSelector not unique, requires index
	CloseMemoriesSelector = `div[role="dialog"] [data-testid="close-button"]`

	CloseSettingsSelector = `div[role="tablist"] [data-testid="close-button"]`
)

var usefulAttributes = []string{"data-testid", "aria-label"}
