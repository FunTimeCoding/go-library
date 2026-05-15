package face

import "github.com/funtimecoding/go-library/pkg/iterm/session"

type ItermSource interface {
	Sessions() ([]session.Session, error)
	CreateTab() (session.Session, error)
	ReadScreen(identifier string) (session.Screen, error)
	ReadHistory(identifier string, count int) (session.History, error)
	SendText(identifier string, text string) error
	SendKey(identifier string, key string) error
	SetTabColor(sessionIdentifier string, red int, green int, blue int) error
	SetTabTitle(tabIdentifier string, title string) error
	MustSessions() []session.Session
	MustCreateTab() session.Session
	MustReadScreen(identifier string) session.Screen
	MustReadHistory(identifier string, count int) session.History
}
