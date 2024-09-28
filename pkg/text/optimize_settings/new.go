package optimize_settings

func New() *Settings {
	return &Settings{
		NewlineAtEnd:      true,
		AllowedBlankLines: 1,
	}
}
