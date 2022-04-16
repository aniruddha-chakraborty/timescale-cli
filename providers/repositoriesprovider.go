package providers

import (
	"github.com/aniruddha-chakraborty/timescale-cli/container"
	"github.com/aniruddha-chakraborty/timescale-cli/database/repositories"
	"github.com/aniruddha-chakraborty/timescale-cli/services"
)

type RepositoryProvider struct {

}

func (r *RepositoryProvider) Register(global_container *container.Container) {

	repository_container := &container.Container{
		Services: make(map[string]interface{}),
	}

	postgres := global_container.Get("dbprovider").(*services.Database)
	postgresConnection := postgres.PostgresConnection()

	repository_container.Set("cpu_usage",&repositories.CpuUsage{ Connection: postgresConnection})
	global_container.Set("repositories",repository_container)
}
