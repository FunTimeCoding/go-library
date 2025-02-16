package web

// noinspection HttpUrlsUsage
const (
	InsecureScheme = "http"
	SecureScheme   = "https"
	SocketScheme = "ws"

	ContentTypeHeader   = "Content-Type"
	UserAgentHeader     = "User-Agent"
	AuthorizationHeader = "Authorization"
	AcceptHeader        = "Accept"
	RealAddressHeader   = "X-Real-Ip"
	RefererHeader       = "Referer"

	FormDataContentType = "multipart/form-data"
	IconContentType     = "image/x-icon"
	ObjectContentType   = "application/json"
	TextContentType     = "text/plain"

	GetMethod   = "GET"
	PostMethod  = "POST"
	PutMethod   = "PUT"
	PatchMethod = "PATCH"

	Localhost = "localhost"

	SecureSchemePrefix = "https://"

	SessionCookie      = "session"
	LastLocationCookie = "last_location"

	ListenAddress  = ":8080"
	MetricsAddress = ":9090"
)

// noinspection HttpUrlsUsage
const InsecureSchemePrefix = "http://"
