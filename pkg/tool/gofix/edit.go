package gofix

import "go/token"

type edit struct {
	position token.Pos
	end      token.Pos
	newText  string
}
