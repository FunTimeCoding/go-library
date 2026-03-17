package cache

import "github.com/anthropics/anthropic-sdk-go"

func ToParameter(m Mode) anthropic.CacheControlEphemeralParam {
	p := anthropic.NewCacheControlEphemeralParam()

	if m == OneHour {
		p.TTL = anthropic.CacheControlEphemeralTTLTTL1h
	}

	return p
}
