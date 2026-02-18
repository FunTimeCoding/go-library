package face

type Component interface {
	Start()
	Stop(verbose bool)
}
