package notation

func Decode(
	value string,
	structure any,
) error {
	return DecodeBytes([]byte(value), structure)
}
