package providers

import (
	"github.com/aniruddha-chakraborty/timescale-cli/container"
	"github.com/aniruddha-chakraborty/timescale-cli/services"
)

type CsvProvider struct {

}

func (csv *CsvProvider) Register(container *container.Container) {
	eventbus := container.Get("eventbus").(*services.EventBus)
	container.Set("csv", &services.CsvHandler{
		EventBus : eventbus,
	})
}