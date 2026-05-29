package service_tester

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gokubernetesd/store/mute_rule"
)

func (t *Tester) MuteRules() []mute_rule.MuteRule {
	result, e := t.Service.MuteRules()
	errors.PanicOnError(e)

	return result
}
