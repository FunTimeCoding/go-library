package browser_tester

import (
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (b *Browser) Evaluate(
	expression string,
	result any,
) {
	b.T.Helper()
	errors.PanicOnError(
		chromedp.Run(b.Context, chromedp.Evaluate(expression, result)),
	)
}
