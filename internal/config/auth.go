package config

// AuthConf holds the configurations for authentication.
type AuthConf struct {
	Endpoint string                `yaml:"endpoint"`
	Clients  map[string]ClientConf `yaml:"clients"`
}

// ClientConf holds the configurations for clients' ids and secrets.
type ClientConf struct {
	ID     string `yaml:"client_id"`
	Secret string `yaml:"client_secret"`
}
