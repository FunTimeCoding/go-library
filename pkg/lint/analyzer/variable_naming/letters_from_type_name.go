package variable_naming

import (
	"go/types"
	"strings"
)

func lettersFromTypeName(t types.Type) []string {
	t = deref(t)
	named, ok := t.(*types.Named)

	if !ok {
		return nil
	}

	name := named.Obj().Name()
	lastWord := extractLastWord(name)

	return lettersFromWord(strings.ToLower(lastWord))
}
