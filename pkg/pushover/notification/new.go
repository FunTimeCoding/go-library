package notification

func New(
	user string,
	token string,
	text string,
) *Notification {
	return &Notification{User: user, Token: token, Message: text}
}
