package providers

import (
	"github.com/aniruddha-chakraborty/timescale-cli/container"
	"github.com/aniruddha-chakraborty/timescale-cli/database/repositories"
	"github.com/aniruddha-chakraborty/timescale-cli/services"
)

type WorkerProvider struct {

}

func (queue *WorkerProvider) Register(global_container *container.Container) {
	calculator := global_container.Get("calculator").(*services.Calculator)
	eventListener := global_container.Get("eventbus").(*services.EventBus)
	waitgroup := global_container.Get("waitgroup").(*services.GlobalWaitGroup)

	repository := global_container.Get("repositories").(*container.Container)
	cpuUsage   := repository.Get("cpu_usage").(*repositories.CpuUsage)

	global_container.Set("workers", &services.Workers{
		Calculator: calculator,
		EventListener: eventListener,
		WaitGroup: waitgroup,
		CpuUsage: cpuUsage,
	})
}