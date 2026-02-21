package face

type Notifier interface {
	Notify(
		f string,
		a ...any,
	)
}
