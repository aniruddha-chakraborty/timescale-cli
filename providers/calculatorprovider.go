package providers

import (
	"github.com/aniruddha-chakraborty/timescale-cli/container"
	"github.com/aniruddha-chakraborty/timescale-cli/services"
)

type Calculator struct {

}

func (calc *Calculator) Register(container *container.Container) {
	output := container.Get("output").(*services.Output)
	container.Set("calculator", &services.Calculator{
		Output: output,
	})
}
