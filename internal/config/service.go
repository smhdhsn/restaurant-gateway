package config

// ServiceConf holds the configurations for services.
type ServiceConf struct {
	Edible struct {
		Address string `yaml:"address"`
	} `yaml:"edible"`
	Order struct {
		Address string `yaml:"address"`
	} `yaml:"order"`
}
