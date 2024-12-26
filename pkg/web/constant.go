package web

// noinspection HttpUrlsUsage
const (
	InsecureScheme = "http"
	SecureScheme   = "https"

	ContentTypeHeader   = "Content-Type"
	UserAgentHeader     = "User-Agent"
	AuthorizationHeader = "Authorization"
	AcceptHeader        = "Accept"
	RealAddressHeader   = "X-Real-Ip"
	RefererHeader       = "Referer"

	TextContentType     = "text/plain"
	ObjectContentType   = "application/json"
	FormDataContentType = "multipart/form-data"

	GetMethod   = "GET"
	PostMethod  = "POST"
	PutMethod   = "PUT"
	PatchMethod = "PATCH"

	Localhost = "localhost"

	SecureSchemePrefix = "https://"

	SessionCookie      = "session"
	LastLocationCookie = "last_location"

	ListenAddress = ":8080"
)

// noinspection HttpUrlsUsage
const InsecureSchemePrefix = "http://"
