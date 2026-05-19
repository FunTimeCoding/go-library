package client

type Client interface {
	VersionsSince(
		since string,
		limit int,
	) []VersionEntry
	SaveImpression(
		content string,
		source string,
	)
}
