package face

type Downloader interface {
	Download(link string) (string, error)
	Begin(link string) Process
}
