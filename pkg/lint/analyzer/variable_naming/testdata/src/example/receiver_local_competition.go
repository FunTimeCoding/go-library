package example

type AppServer struct{}

func (h *AppServer) Process() { // want `variable h of type \*example.AppServer should be named s`
	var s string // want `variable s of type string should be named t`
	_ = h
	_ = s
}
