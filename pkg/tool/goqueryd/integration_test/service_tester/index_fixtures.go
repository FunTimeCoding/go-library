package service_tester

import (
	"github.com/funtimecoding/go-library/pkg/assert/fixture"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	goqueryd "github.com/funtimecoding/go-library/pkg/tool/goqueryd/constant"
)

func (t *Tester) IndexFixtures() {
	t.Service.AddCollection(
		"test",
		fixture.Path(constant.SearchPath),
		goqueryd.DefaultGlob,
	)
	t.Service.Index("test")
}
