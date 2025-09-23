package errors

func PanicFlush(c Flusher) {
	PanicOnError(c.Flush())
}
