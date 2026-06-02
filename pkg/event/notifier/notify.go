package notifier

func (n *Notifier) Notify() {
	n.mutex.Lock()
	defer n.mutex.Unlock()

	for c := range n.subscribers {
		select {
		case c <- struct{}{}:
		default:
		}
	}
}
