package providers

import (
	"github.com/aniruddha-chakraborty/timescale-cli/container"
	"github.com/aniruddha-chakraborty/timescale-cli/interfaces"
	"github.com/aniruddha-chakraborty/timescale-cli/services"
)

type EventsProvider struct {

}

func (e *EventsProvider) Register(container *container.Container) {
	container.Set("eventbus",&services.EventBus{
		Listeners: make(map[int][]func(structure *interfaces.CsvStructure)),
	})
}
