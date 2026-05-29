package inventory

type Instance struct {
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Secure   bool   `yaml:"secure"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}
