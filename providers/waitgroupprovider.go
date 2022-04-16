package providers

import (
	"github.com/aniruddha-chakraborty/timescale-cli/container"
	"github.com/aniruddha-chakraborty/timescale-cli/services"
)

type WaitGroupProvider struct {

}

func (wait *WaitGroupProvider) Register(container *container.Container) {
	container.Set("waitgroup", &services.GlobalWaitGroup{})
}

