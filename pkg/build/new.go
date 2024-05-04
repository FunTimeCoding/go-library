package build

func New(
	version string,
	gitHash string,
	date string,
) *Build {
	if version == "" {
		version = DefaultVersion
	}

	if gitHash == "" {
		gitHash = DefaultGitHash
	}

	if date == "" {
		date = DefaultDate
	}

	return &Build{
		Version:   version,
		GitHash:   gitHash,
		BuildDate: date,
	}
}
