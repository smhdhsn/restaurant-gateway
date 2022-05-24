package config

// ServerConf holds the configurations for server.
type ServerConf struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
