package face

type Notifier interface {
	Notify()
	Subscribe() chan struct{}
	Unsubscribe(c chan struct{})
}
