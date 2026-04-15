package call_format

import "go/token"

// Fix represents a text replacement for fixing call formatting.
type Fix struct {
	Position token.Pos
	End      token.Pos
	NewText  string
}
