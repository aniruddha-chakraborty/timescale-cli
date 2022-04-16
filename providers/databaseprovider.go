package providers

import (
	"github.com/aniruddha-chakraborty/timescale-cli/container"
	"github.com/aniruddha-chakraborty/timescale-cli/services"
)

type DatabaseProvider struct {

}

func (db *DatabaseProvider) Register(container *container.Container) {
	config := container.Get("config").(*services.Config)
	container.Set("dbprovider", &services.Database{
		Config: config,
	})
}

