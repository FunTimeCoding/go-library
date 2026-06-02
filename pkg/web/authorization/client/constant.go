package client

const (
	ClientIdentifierEnvironment = "GATE_CLIENT_IDENTIFIER"
	ClientSecretEnvironment     = "GATE_CLIENT_SECRET"
	IssuerEnvironment           = "GATE_ISSUER"
	SecretEnvironment           = "GATE_SECRET"
	CallbackLocatorEnvironment  = "GATE_CALLBACK_LOCATOR"
	SignInPathEnvironment       = "GATE_SIGN_IN_PATH"
)

const (
	flowCookie    = "authorization_flow"
	subjectCookie = "authorization_subject"
	defaultScope  = "openid offline"
)
