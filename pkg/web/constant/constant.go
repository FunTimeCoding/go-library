package constant

import "net/http"

// noinspection HttpUrlsUsage
const (
	InsecureScheme = "http"
	SecureScheme   = "https"
	SocketScheme   = "ws"

	ContentTypeHeader   = "Content-Type"
	UserAgentHeader     = "User-Agent"
	AuthorizationHeader = "Authorization"
	AcceptHeader        = "Accept"
	RealAddressHeader   = "X-Real-Ip"
	RefererHeader       = "Referer"

	FormDataContentType = "multipart/form-data"
	IconContentType     = "image/x-icon"
	MarkdownContentType = "text/markdown"
	ObjectContentType   = "application/json"
	TextContentType     = "text/plain"

	DeleteMethod   = "DELETE"
	GetMethod      = "GET"
	PatchMethod    = "PATCH"
	PostMethod     = "POST"
	PropfindMethod = "PROPFIND"
	PutMethod      = "PUT"

	Localhost = "localhost"

	SecureSchemePrefix = "https://"

	SessionCookie      = "session"
	LastLocationCookie = "last_location"

	ListenAddress  = ":8080"
	MetricsAddress = ":9090"

	BasicPrefix  = "Basic"
	BearerPrefix = "Bearer"
)

// noinspection HttpUrlsUsage
const InsecureSchemePrefix = "http://"

var OkayStatusCodes = []int{http.StatusOK, http.StatusCreated}
