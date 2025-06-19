package site

const (
	// NewChatSelector not unique, requires index
	NewChatSelector = `a[data-testid="create-new-chat-button"]`

	ProfileButtonSelector = `[data-testid="profile-button"]`

	SettingsButtonSelector = `[data-testid="settings-menu-item"]`

	PersonalizationTabSelector = `[data-testid="personalization-tab"]`

	ManageMemoriesSelector = `//a[text()="Manage memories"]`

	// CloseDialogSelector not unique, requires index
	CloseDialogSelector = `div[role="dialog"] [data-testid="close-button"]`

	CloseSettingsSelector = `div[role="tablist"] [data-testid="close-button"]`
)

var usefulAttributes = []string{"data-testid", "aria-label"}
