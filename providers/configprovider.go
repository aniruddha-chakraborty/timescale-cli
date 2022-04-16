package providers

import (
	"github.com/aniruddha-chakraborty/timescale-cli/container"
	"github.com/aniruddha-chakraborty/timescale-cli/services"
)

type ConfigProvider struct {

}

func (c *ConfigProvider) Register(container *container.Container) {
	container.Set("config", &services.Config{Configuration: make(map[string]map[string]string)})
	var config = container.Get("config").(*services.Config)
	config.Load()
}
