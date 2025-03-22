package amtool

type Option struct {
	Locator          string `yaml:"alertmanager.url"`
	WebConfiguration string `yaml:"http.config.file"`
}
