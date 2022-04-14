package main

import (
	"flag"
	"fmt"
	"github.com/aniruddha-chakraborty/timescale-cli/container"
	"github.com/aniruddha-chakraborty/timescale-cli/providers"
	"github.com/aniruddha-chakraborty/timescale-cli/services"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	dir, _ := os.Getwd()
	envdirectory := fmt.Sprintf("%s/.env",dir)
	err := godotenv.Load(envdirectory)
	if err != nil {
		log.Fatal(err)
	}
	containers := &container.Container {
		Services: make(map[string]interface{}),
	}

	configprovider 			:= &providers.ConfigProvider{}
	queueprovider 			:= &providers.QueueProvider{}
	csvprovider 			:= &providers.CsvProvider{}

	configprovider.Register(containers)
	queueprovider.Register(containers)
	csvprovider.Register(containers)

	//configs	:= containers.Get("config").(*services.Config)
	filePtr := flag.String("file", "foo", "a string")
	flag.Parse()
	//fmt.Println("word:", *filePtr)

	csv := containers.Get("csv").(*services.CsvHandler)
	csv.Read(*filePtr)
	//println(configs.Get("TIMESCALE","HOST"))
}