package debian

import (
	"github.com/coreos/go-semver/semver"
	"log"
)

func Codename(v *semver.Version) string {
	switch v.Major {
	case Bookworm.Major:
		return Bookworm.Codename
	default:
		log.Panicf("unsupported major version %d", v.Major)
	}

	return ""
}
