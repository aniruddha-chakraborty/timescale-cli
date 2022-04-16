package interfaces

type CsvStructure struct {
	Hostname string `csv:"hostname"`
	StartTime string `csv:"start_time"`
	EndTime string `csv:"end_time"`
}

type QueueData struct {
	Host string
	StartTime string
	EndTime string
}
