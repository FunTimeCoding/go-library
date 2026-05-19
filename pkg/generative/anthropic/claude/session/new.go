package session

func New(identifier string) *Session {
	return &Session{Identifier: identifier}
}
