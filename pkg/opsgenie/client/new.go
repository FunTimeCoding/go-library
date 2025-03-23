package client

func New(key string) *Client {
	l := Configuration(key)

	return &Client{
		Account:      Account(l),
		Alert:        Alert(l),
		Contact:      Contact(l),
		Deployment:   Deployment(l),
		Escalation:   Escalation(l),
		Forward:      Forward(l),
		Heartbeat:    Heartbeat(l),
		Incident:     Incident(l),
		Integration:  Integration(l),
		Log:          Log(l),
		Maintenance:  Maintenance(l),
		Notification: Notification(l),
		Policy:       Policy(l),
		Schedule:     Schedule(l),
		Service:      Service(l),
		Team:         Team(l),
		User:         User(l),
	}
}
