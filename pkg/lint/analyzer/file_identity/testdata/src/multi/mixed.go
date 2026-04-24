package multi

type Service struct{}

func Helper() {} // want `multiple identities in one file: Service and Helper`
