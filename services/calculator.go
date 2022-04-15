package services

import (
	"fmt"
	"sync"
)

type Calculator struct {
	queriesProcessed int64
	totalProcessing float64
	minQueryTime float64
	medianQueryTime float64
	avgQueryTime float64
	maxQueryTime float64
	su sync.Mutex
}

func (c *Calculator) Calculate(queryTime float64) {
	c.su.Lock()
	c.queriesProcessed = c.queriesProcessed + 1
	c.totalProcessing  = c.totalProcessing + queryTime
	if c.minQueryTime > queryTime {
		c.minQueryTime = queryTime
	}
	if c.maxQueryTime < queryTime {
		c.maxQueryTime = queryTime
	}
	c.avgQueryTime = float64(c.queriesProcessed) / c.totalProcessing
	c.medianQueryTime = c.totalProcessing / 2
	c.su.Unlock()
	//fmt.Fprintf(c.Output,"total Query Processed %d",c.queriesProcessed)
	//fmt.Fprintf(c.Output, "Total processing time %f seconds", (c.totalProcessing/1000)/1000)
	//fmt.Fprintf(c.Output,"Minimum Query time %f Seconds",c.minQueryTime)
	//fmt.Fprintf(c.Output,"Maximum Query Time %f Seconds",c.maxQueryTime)
}

func (c *Calculator) PrintResult() {
	fmt.Println("total Query Processed ",c.queriesProcessed)
	fmt.Println("Total processing time", (c.totalProcessing/1000)/1000)
	fmt.Println("Minimum Query time ",(c.minQueryTime/1000)/1000)
	fmt.Println("Maximum Query Time ",(c.maxQueryTime/1000)/1000)
}