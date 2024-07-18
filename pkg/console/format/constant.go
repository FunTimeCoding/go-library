package format

var (
	Plain = New()
	Color = New().Color()

	Raw      = New().Raw()
	RawColor = New().Raw().Color()

	Extended      = New().Extended()
	ExtendedColor = New().Extended().Color()
)
