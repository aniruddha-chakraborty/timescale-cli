package services

import (
	"sync"
	"time"
)

type Calculator struct {
	queriesProcessed int64
	totalProcessing int64
	minQueryTime int64
	medianQueryTime int64
	avgQueryTime int64
	maxQueryTime int64
	su sync.Mutex
}

func (c *Calculator) Calculate(queryTime int64,startTime time.Time,totalInQueue int64) {
	newtime := time.Now().Sub(startTime)
	newtime.Nanoseconds()
	c.su.Lock()
	c.queriesProcessed = c.queriesProcessed + 1
	c.totalProcessing  = c.totalProcessing + newtime.Nanoseconds()
	if c.minQueryTime > queryTime {
		c.minQueryTime = queryTime
	}
	if c.maxQueryTime > queryTime {
		c.maxQueryTime = queryTime
	}
	c.avgQueryTime = c.queriesProcessed / c.totalProcessing
	c.medianQueryTime = c.totalProcessing / 2
	c.su.Unlock()

}