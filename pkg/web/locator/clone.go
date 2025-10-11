package locator

import "net/url"

func (l *Locator) Clone() *Locator {
	values := url.Values{}

	for k, v := range l.value {
		values[k] = v
	}

	return &Locator{
		value:    values,
		scheme:   l.scheme,
		host:     l.host,
		port:     l.port,
		basePath: l.basePath,
		path:     l.path,
	}
}
