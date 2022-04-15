package services

import (
	"fmt"
	"github.com/aniruddha-chakraborty/timescale-cli/interfaces"
	"github.com/gocarina/gocsv"
	"os"
)

type CsvHandler struct {
	Queue *Queue
}

func (c *CsvHandler) Read(filename string) []*interfaces.CsvStructure {
	clientsFile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer clientsFile.Close()
	var clients []*interfaces.CsvStructure

	if err := gocsv.UnmarshalFile(clientsFile, &clients); err != nil { // Load clients from file
		panic(err)
	}

	for _, client := range clients {
		println(client.Hostname)
		//c.Queue.Insert(&interfaces.QueueData{
		//	Host: client.Hostname,
		//	StartTime: client.StartTime,
		//	EndTime: client.EndTime,
		//})
	}
	fmt.Println("test")
	return clients
}