package notation

func DecodeAny(
	raw any,
	into any,
) {
	DecodeStrict(Encode(raw, false), into, false)
}
