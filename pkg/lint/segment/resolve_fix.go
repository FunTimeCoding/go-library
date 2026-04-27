package segment

import (
	"go/token"
	"go/types"
)

// ResolveFix picks a replacement for the matched segment, walking only
// parent scopes for collision detection. Used by the analyzer where the
// types.Info may not have visibility into all child scopes (e.g. test
// files loaded in a separate pass).
func ResolveFix(
	name string,
	segment string,
	applicable []string,
	o types.Object,
) string {
	scope := o.Parent()
	fix := ChooseFix(name, applicable)
	replacement := ReplaceSegment(name, segment, fix)

	if token.IsKeyword(replacement) {
		return ""
	}

	if scope == nil || scope.Lookup(replacement) == nil {
		return fix
	}

	if len(fix) > 1 {
		return ""
	}

	multiCharacter := FirstMultiCharacter(applicable)

	if multiCharacter == "" {
		return ""
	}

	for _, r := range multiCharacter {
		letter := string(r)
		candidate := ReplaceSegment(name, segment, letter)

		if scope.Lookup(candidate) == nil {
			return letter
		}
	}

	return ""
}
