package gomemoryd

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/service"
)

func reconcileMemories(svc *service.Service) {
	memories, e := svc.Store.ListMemories("", "", true)
	errors.PanicOnError(e)

	for _, m := range memories {
		full, f := svc.Store.GetMemory(m.Identifier)
		errors.PanicOnError(f)
		svc.MustIndexMemory(
			fmt.Sprintf("memory/%d", full.Identifier),
			full.Content,
		)
	}
}
