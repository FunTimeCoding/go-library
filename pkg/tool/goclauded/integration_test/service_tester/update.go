package service_tester

import "github.com/funtimecoding/go-library/pkg/assert"

func (o *Tester) Update(
	sessionIdentifier string,
	name string,
	topic string,
	files string,
) {
	assert.FatalOnError(
		o.t,
		o.Service.Update(sessionIdentifier, name, topic, files),
	)
}
