package mock_notifier

func New(prefix string) *Notifier {
	return &Notifier{prefix: prefix}
}
