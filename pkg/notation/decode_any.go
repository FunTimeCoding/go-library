package notation

func DecodeAny(
	raw any,
	into any,
) {
	MustDecode(Encode(raw, false), into, false)
}
