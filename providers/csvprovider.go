package providers

import (
	"github.com/aniruddha-chakraborty/timescale-cli/container"
	"github.com/aniruddha-chakraborty/timescale-cli/services"
)

type CsvProvider struct {

}

func (csv *CsvProvider) Register(container *container.Container) {
	queue := container.Get("queue").(*services.Queue)
	container.Set("csv", &services.CsvHandler{
		Queue: queue,
	})
}