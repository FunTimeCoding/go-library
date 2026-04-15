package lint

import (
	"path/filepath"
	"strings"
)

func isKnownBinaryExtension(path string) bool {
	extension := strings.ToLower(filepath.Ext(path))

	switch extension {
	case ".png", ".jpg", ".jpeg", ".gif", ".ico", ".webp", ".bmp", ".tiff":
		return true
	case ".svg":
		return true
	case ".pdf":
		return true
	case ".woff", ".woff2", ".ttf", ".otf", ".eot":
		return true
	case ".zip", ".gz", ".tar", ".bz2", ".xz", ".zst":
		return true
	case ".mp3", ".mp4", ".wav", ".ogg", ".webm", ".flac":
		return true
	}

	return false
}
