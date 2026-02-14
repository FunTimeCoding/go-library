package tab

type Tab struct {
	Description    string `json:"description"`
	Frontend       string `json:"devtoolsFrontendUrl"`
	Id             string `json:"id"`
	Title          string `json:"title"`
	Type           string `json:"type"`
	Locator        string `json:"url"`
	WebSocketDebug string `json:"webSocketDebuggerUrl"`
	FaviconUrl     string `json:"faviconUrl,omitempty"`
}
