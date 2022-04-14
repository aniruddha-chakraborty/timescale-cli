package services

import "sync"

type Calculator struct {
	queriesProcessed float32
	totalProcessing float32
	minQueryTime float64
	medianQueryTime float32
	avgQueryTime float32
	maxQueryTime float32
	totalQueries int32
	su sync.Mutex
}

func (c *Calculator) Calculate(queryTime int64) {
	c.su.Lock()
	c.totalQueries = c.totalQueries + 1
	c.su.Unlock()
}
