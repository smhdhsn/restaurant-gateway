package config

// This section holds names of services which will be used withing this service as string constant.
const (
	EdibleService = "edible"
	OrderService  = "order"
)

// ServiceConf holds the configurations for service.
type ServiceConf struct {
	Address string `yaml:"address"`
}
