package object_notation

func MarshallIndent(a any) string {
	return string(MarshallIndentBytes(a))
}
