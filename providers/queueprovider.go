package providers

import (
	"github.com/aniruddha-chakraborty/timescale-cli/container"
	"github.com/aniruddha-chakraborty/timescale-cli/services"
)

type QueueProvider struct {
	
}

func (queue *QueueProvider) Register(container *container.Container) {
	container.Set("queue", &services.Queue{})
}
