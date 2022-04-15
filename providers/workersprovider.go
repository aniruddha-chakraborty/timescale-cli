package providers

import (
	"github.com/aniruddha-chakraborty/timescale-cli/container"
	"github.com/aniruddha-chakraborty/timescale-cli/services"
)

type WorkerProvider struct {

}

func (queue *WorkerProvider) Register(container *container.Container) {
	calculator := container.Get("calculator").(*services.Calculator)
	eventListener := container.Get("eventbus").(*services.EventBus)

	container.Set("workers", &services.Workers{
		Calculator: calculator,
		EventListener: eventListener,
	})
}