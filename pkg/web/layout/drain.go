package layout

func drain(c chan struct{}) {
	for {
		select {
		case <-c:
		default:
			return
		}
	}
}
