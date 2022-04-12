package services

import "os"

type Config struct {
	Configuration map[string]map[string]string
}

func (c *Config) Ifelse(data1 string , data2 string,output string,output_else string ) string {
	if data1 == data2 {
		return output
	}
	return output_else
}

func (c *Config) Load() {

	timescaledb 		:= make(map[string]string)

	timescaledb["HOST"] 		= c.Ifelse(os.Getenv("TIMESCALEHOST"),"","",os.Getenv("TIMESCALEHOST"))
	timescaledb["USER"] 		= c.Ifelse(os.Getenv("TIMESCALEUSER"),"","",os.Getenv("TIMESCALEUSER"))
	timescaledb["PASSWORD"] 	= c.Ifelse(os.Getenv("TIMESCALEPASSWORD"),"","",os.Getenv("TIMESCALEPASSWORD"))
	timescaledb["DATABASE"] 	= c.Ifelse(os.Getenv("TIMESCALEDATABASE"),"","",os.Getenv("TIMESCALEDATABASE"))
	timescaledb["PORT"] 		= c.Ifelse(os.Getenv("TIMESCALEPORT"),"","",os.Getenv("TIMESCALEPORT"))


}

func (c *Config) Get(root string , config string) string {
	return c.Configuration[root][config]
}

func (c *Config) Set(root string, config string,value string) string {
	c.Configuration[root] = map[string]string{
		config: value,
	}
	return c.Configuration[root][config]
}