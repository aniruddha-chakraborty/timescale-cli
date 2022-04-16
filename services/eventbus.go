package services

import "github.com/aniruddha-chakraborty/timescale-cli/interfaces"

type convert func(structure *interfaces.CsvStructure)

type EventBus struct {
	Listeners map[int][]func(structure *interfaces.CsvStructure)
}

func (e *EventBus) AddListener(event int , fn convert ) {
	if _,ok := e.Listeners[event]; ok {
		e.Listeners[event] = append(e.Listeners[event],fn)
	} else {
		e.Listeners[event] = []func(structure *interfaces.CsvStructure){fn}
	}
}

func (e *EventBus) Emit(event int,data *interfaces.CsvStructure) {
	for _,fn := range e.Listeners[event] {
		fn(data)
	}
}