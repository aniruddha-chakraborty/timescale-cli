package main

import (
	"fmt"
	"github.com/aniruddha-chakraborty/timescale-cli/container"
	"github.com/aniruddha-chakraborty/timescale-cli/providers"
	"github.com/aniruddha-chakraborty/timescale-cli/services"
	"github.com/joho/godotenv"
	"log"
	"os"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	dir, _ := os.Getwd()
	envdirectory := fmt.Sprintf("%s/.env",dir)
	err := godotenv.Load(envdirectory)
	if err != nil {
		log.Fatal(err)
	}
	containers := &container.Container{
		Services: make(map[string]interface{}),
	}

	configprovider 			:= &providers.ConfigProvider{}
	configprovider.Register(containers)

	configs	:= containers.Get("config").(*services.Config)
	println(configs.Get("TIMESCALE","HOST"))
}