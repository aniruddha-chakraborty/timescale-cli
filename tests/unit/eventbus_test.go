package unit

import (
	"github.com/aniruddha-chakraborty/timescale-cli/interfaces"
	"github.com/aniruddha-chakraborty/timescale-cli/services"
	"testing"
)

func TestEventBus(t *testing.T) {
	eventbus := services.EventBus{
		Listeners: make(map[int][]func(structure *interfaces.CsvStructure)),
	}

	data := &interfaces.CsvStructure{
		Hostname: "host00001",
		StartTime: "2017-01-01 08:59:22",
		EndTime: "2017-01-01 09:59:22",
	}

	eventbus.AddListener(1,func(structure *interfaces.CsvStructure) {
		if structure != data {
			t.Error("Event listeners not working")
		}
	})

	eventbus.Emit(1,data)
}
