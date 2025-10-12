package locator

import "net/url"

func (l *Locator) Copy() *Locator {
	value := url.Values{}

	for k, v := range l.value {
		value[k] = v
	}

	fragmentValue := url.Values{}

	for k, v := range l.fragmentValue {
		fragmentValue[k] = v
	}

	return &Locator{
		scheme:        l.scheme,
		user:          l.user,
		password:      l.password,
		host:          l.host,
		port:          l.port,
		basePath:      l.basePath,
		path:          l.path,
		trail:         l.trail,
		value:         value,
		fragment:      l.fragment,
		fragmentValue: fragmentValue,
	}
}
