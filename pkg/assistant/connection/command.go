package connection

type command interface {
	SetIdentifier(uint64)
}
