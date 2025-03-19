package label

func New(
	k string,
	v string,
) *Label {
	return &Label{Key: k, Value: v}
}
