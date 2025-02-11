package server

type Server struct {
	Logf func(
		f string,
		v ...any,
	)
}
