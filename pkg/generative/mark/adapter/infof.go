package adapter

func (a *Adapter) Infof(
	format string,
	v ...any,
) {
	a.Logger.Plain(format, v...)
}
