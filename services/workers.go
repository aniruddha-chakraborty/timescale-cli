package services

import (
	"github.com/aniruddha-chakraborty/timescale-cli/interfaces"
	"time"
)

type Workers struct {
	EventListener *EventBus
	Calculator *Calculator
}

func (w *Workers) Listen() {
	w.EventListener.AddListener(1,func(structure *interfaces.CsvStructure) {
		//fmt.Println("testing")
		go w.exec(structure)
	})

	w.EventListener.AddListener(2,func(structure *interfaces.CsvStructure) {
		go w.exec(structure)
	})
}

func ( w *Workers ) exec(data *interfaces.CsvStructure) {
	//fmt.Println("testing")
	start := time.Now()
	time.Sleep(1 * time.Microsecond)
	end   := time.Now()



	d := float64(end.UnixNano() - start.UnixNano())
	//fmt.Println(d)
	w.Calculator.Calculate(d)
}