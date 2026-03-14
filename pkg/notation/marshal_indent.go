package notation

func MarshalIndent(a any) string {
	return string(MarshalIndentBytes(a))
}
