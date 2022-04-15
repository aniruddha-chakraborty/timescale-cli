package services

import "github.com/aniruddha-chakraborty/timescale-cli/interfaces"

type Events struct {
	Listeners map[string][]chan []interfaces.CsvStructure
}

func (e *Events) AddListener(event string, channel chan []interfaces.CsvStructure ) {
	if _,ok := e.Listeners[event]; ok {
		e.Listeners[event] = append(e.Listeners[event],channel)
	} else {
		e.Listeners[event] = []chan []interfaces.CsvStructure{channel}
	}
}

func (e *Events) RemoveListener(event string, channel chan []interfaces.CsvStructure) {
	if _,ok := e.Listeners[event]; ok {
		for i := range e.Listeners[event] {
			if e.Listeners[event][i] == channel {
				e.Listeners[event][i] = e.Listeners[event][len(e.Listeners[event])-1]
				e.Listeners[event] =  e.Listeners[event][:len(e.Listeners[event])-1]
			}
		}
	}
}

func (e *Events) Emit(event string,data []interfaces.CsvStructure) {
	if _, ok := e.Listeners[event]; ok {
		for _, handler := range e.Listeners[event] {
			//handler <- data
			go func(handler chan []interfaces.CsvStructure) {
				handler <- data
			}(handler)
		}
	}
}