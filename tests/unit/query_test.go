package unit

import (
	"github.com/aniruddha-chakraborty/timescale-cli/database/repositories"
	"github.com/aniruddha-chakraborty/timescale-cli/interfaces"
	"github.com/aniruddha-chakraborty/timescale-cli/services"
	"testing"
)

func TestQuery(t *testing.T) {
	database := services.Database{}
	database.PostgresConnect()

	repository := &repositories.CpuUsage{
		Connection: database.PostgresConnection(),
	}

	data := &interfaces.CsvStructure{
		Hostname: "host_000008",
		StartTime: "2017-01-01 08:59:22",
		EndTime: "2017-01-01 09:59:22",
	}
	output := repository.GetSelected(data.StartTime,data.EndTime,data.Hostname)
	if len(output) <= 0 {
		t.Error("Query not working")
	}
}
