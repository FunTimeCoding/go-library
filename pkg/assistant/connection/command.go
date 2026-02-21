package connection

type command interface {
	setIdentifier(uint64)
}
