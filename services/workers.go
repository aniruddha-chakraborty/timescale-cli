package services

import (
	"github.com/aniruddha-chakraborty/timescale-cli/database/repositories"
	"github.com/aniruddha-chakraborty/timescale-cli/interfaces"
	"time"
)

type Workers struct {
	EventListener *EventBus
	Calculator *Calculator
	WaitGroup *GlobalWaitGroup
	CpuUsage *repositories.CpuUsage
}

func (w *Workers) Listen() {
	totalWorkers := 2
	for i := 0; i < totalWorkers; i++ {
		w.EventListener.AddListener(i,func(structure *interfaces.CsvStructure) {
			w.WaitGroup.Instance().Add(1)
			go w.exec(structure)
		})
	}
}

func ( w *Workers ) exec(data *interfaces.CsvStructure) {
	start := time.Now()
	w.CpuUsage.GetSelected(data.StartTime,data.EndTime,data.Hostname)
	end   := time.Now()

	d := float64(end.UnixNano() - start.UnixNano())
	w.Calculator.Calculate(d)
	w.WaitGroup.Instance().Done()
}