package component

func (w *Component) Stop(_ bool) {
	w.Worker.Stop()
}
