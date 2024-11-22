package tunnel

func New() *Tunnel {
	return &Tunnel{
		started:   make(chan struct{}),
		listening: make(chan struct{}),
		stopped:   make(chan struct{}),
		CleanStop: true,
	}
}
