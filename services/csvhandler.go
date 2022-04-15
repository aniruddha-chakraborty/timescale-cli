package services

import (
	"fmt"
	"github.com/aniruddha-chakraborty/timescale-cli/interfaces"
	"github.com/gocarina/gocsv"
	"os"
	"strconv"
)

type CsvHandler struct {
	EventBus *EventBus
	Calculator *Calculator
}

func (c *CsvHandler) Read(filename string) {
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
		hostId := client.Hostname[5:len(client.Hostname)]
		n, err := strconv.ParseInt(hostId, 10, 64)
		if err != nil {
			fmt.Printf("%d of type %T", n, n)
		}
		//fmt.Println(hostId)
		c.EventBus.Emit(int(n),client)
	}

	c.Calculator.PrintResult()
}