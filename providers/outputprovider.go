package providers

import (
	"github.com/aniruddha-chakraborty/timescale-cli/container"
	"github.com/aniruddha-chakraborty/timescale-cli/services"
)

type OutputProvider struct {

}

func (out *OutputProvider) Register(container *container.Container) {
	container.Set("output", &services.Output{})
}
