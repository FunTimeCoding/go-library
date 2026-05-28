package environment

func (e *Environment) Load(path string) error {
	overlay, loadError := eval(path, e.base)

	if loadError != nil {
		return loadError
	}

	e.mutex.Lock()
	e.overlay = overlay
	e.mutex.Unlock()

	return nil
}
