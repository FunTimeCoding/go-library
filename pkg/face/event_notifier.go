package face

type EventNotifier interface {
	Notify()
	Subscribe() chan struct{}
	Unsubscribe(c chan struct{})
}
