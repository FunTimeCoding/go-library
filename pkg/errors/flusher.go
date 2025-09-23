package errors

type Flusher interface {
	Flush() error
}
