package amtool

type Configuration struct {
	Locator          string `yaml:"alertmanager.url"`
	WebConfiguration string `yaml:"http.config.file"`
}
