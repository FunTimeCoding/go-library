package lint

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/lint/concern"
	"github.com/funtimecoding/go-library/pkg/lint/constant"
)

func concernMessage(c *concern.Concern) string {
	switch c.Key {
	case constant.BlankInsideFunctionKey,
		constant.ExtraneousBlankLineKey,
		constant.ExtraneousTopLevelBlankKey:

		return fmt.Sprintf("removed blank line (line %d)", c.Line)
	case constant.MissingBlankBeforeControlKey,
		constant.MissingBlankAfterControlKey,
		constant.MissingBlankBeforeExitKey,
		constant.MissingBlankBeforeDeclarationKey,
		constant.MissingBlankBetweenVariableConstantKey:

		return fmt.Sprintf("inserted blank line (line %d)", c.Line)
	case constant.ImportBlankKey:
		return "removed import blank line"
	case constant.UnsortedImportsKey:
		return "sorted imports"
	case constant.SingleMultiImportKey:
		return "collapsed single multi-import"
	case constant.EmptyFunctionBodyKey:
		return "cleaned empty function body"
	case constant.VariableGroupingKey:
		return "grouped var declarations"
	default:
		return c.Text
	}
}
