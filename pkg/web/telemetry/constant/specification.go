package constant

// https://opentelemetry.io/docs/specs/semconv/http/http-spans/

// Required attributes
const (
	RequestMethod = "http.request.method"
	Path          = "url.path"
	Scheme        = "url.scheme"
)

// Conditionally required
const (
	Query  = "url.query"
	Status = "http.response.status_code"
)

// Recommended
const (
	Client    = "client.address"
	Peer      = "network.peer.address"
	Protocol  = "network.protocol.version"
	Server    = "server.address"
	UserAgent = "user_agent.original"
)

// Opt-in attributes for webhooks
const (
	BodySize     = "http.request.body.size"
	Body         = "http.request.body"
	HeaderPrefix = "http.request.header"
)
