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
	waitgroupProvider 		:= &providers.WaitGroupProvider{}
	eventbusprovider 		:= &providers.EventsProvider{}
	workersprovider 		:= &providers.WorkerProvider{}
	dbprovider				:= &providers.DatabaseProvider{}
	repositoriesProvider 	:= &providers.RepositoryProvider{}
	outputprovider 			:= &providers.OutputProvider{}
	calculatorprovider 		:= &providers.Calculator{}
	csvprovider 			:= &providers.CsvProvider{}

	configprovider.Register(containers)
	waitgroupProvider.Register(containers)
	outputprovider.Register(containers)
	//output := containers.Get("output").(*services.Output)
	//output.Init()
	calculatorprovider.Register(containers)
	eventbusprovider.Register(containers)
	workersprovider.Register(containers)
	dbprovider.Register(containers)
	repositoriesProvider.Register(containers)
	csvprovider.Register(containers)


	workers := containers.Get("workers").(*services.Workers)
	workers.Listen()

	filePtr := flag.String("file", "foo", "a string")
	flag.Parse()
	csv := containers.Get("csv").(*services.CsvHandler)
	csv.Read(*filePtr)
}