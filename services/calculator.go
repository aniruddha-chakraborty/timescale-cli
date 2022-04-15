package services

import (
	"fmt"
	"sync"
	"time"
)

type Calculator struct {
	queriesProcessed int64
	totalProcessing float64
	minQueryTime float64
	medianQueryTime float64
	avgQueryTime float64
	maxQueryTime float64
	Output *Output
	su sync.Mutex
}

func (c *Calculator) Calculate(queryTime float64,startTime time.Time,totalInQueue int64) {
	newtime := time.Now().Sub(startTime)
	newtime.Nanoseconds()
	c.su.Lock()
	c.queriesProcessed = c.queriesProcessed + 1
	c.totalProcessing  = c.totalProcessing + float64(newtime.Nanoseconds())
	if c.minQueryTime > queryTime {
		c.minQueryTime = queryTime
	}
	if c.maxQueryTime > queryTime {
		c.maxQueryTime = queryTime
	}
	c.avgQueryTime = float64(c.queriesProcessed) / c.totalProcessing
	c.medianQueryTime = c.totalProcessing / 2
	c.su.Unlock()

	fmt.Fprintf(c.Output.Writer(),"total Query Processed %d",c.queriesProcessed)
	fmt.Fprintf(c.Output.Writer(), "Total processing time %f seconds", (c.totalProcessing/1000)/1000)
	fmt.Fprintf(c.Output.Writer(),"Minimum Query time %f Seconds",c.minQueryTime)
	fmt.Fprintf(c.Output.Writer(),"Maximum Query Time %f Seconds",c.maxQueryTime)
	c.Output.Writer().Print()
	c.Output.Writer().Clear()
}