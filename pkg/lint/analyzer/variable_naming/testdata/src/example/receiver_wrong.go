package example

type ReceiverTarget struct{}

func (h *ReceiverTarget) DoThing() { // want `variable h of type \*example.ReceiverTarget should be named t`
	_ = h
}
