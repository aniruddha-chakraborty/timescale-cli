package services

import (
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
		c.Queue.Insert(client)
	}

	return clients
}