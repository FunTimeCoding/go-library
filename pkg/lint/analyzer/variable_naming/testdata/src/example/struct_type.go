package example

type MyClient struct{}

func StructType() {
	x := &MyClient{} // want `variable x of type \*example\.MyClient should be named c`
	_ = x
}
