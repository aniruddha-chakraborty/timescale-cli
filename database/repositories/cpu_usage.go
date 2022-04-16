package repositories

import (
	"github.com/aniruddha-chakraborty/timescale-cli/database/models"
	"github.com/jinzhu/gorm"
)

type CpuUsage struct {
	Connection *gorm.DB
}


func (c *CpuUsage) GetSelected(startTime string , endTime string, host string) []models.CpuUsage {
	var cpuUsage []models.CpuUsage
	c.Connection.Raw("select host, avg(usage), date_trunc('minute',ts) as ts FROM cpu_usage WHERE ts BETWEEN ? and ? AND host = ? GROUP BY host,date_trunc('minute',ts) ORDER BY ts;", startTime,endTime,host).Scan(&cpuUsage)
	return cpuUsage
}
