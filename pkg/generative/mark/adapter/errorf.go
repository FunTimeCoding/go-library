package adapter

func (a *Adapter) Errorf(
	format string,
	v ...any,
) {
	a.Logger.Plain(format, v...)
}
