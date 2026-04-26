package gofix

func splice(
	b []byte,
	offset int,
	length int,
	replacement []byte,
) []byte {
	var result []byte
	result = append(result, b[:offset]...)
	result = append(result, replacement...)
	result = append(result, b[offset+length:]...)

	return result
}
