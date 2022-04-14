package services

type Calculator struct {
	queriesProcessed float32
	totalProcessing float32
	minQueryTime float64
	medianQueryTime float32
	avgQueryTime float32
	maxQueryTime float32
	totalQueries int32
}

func (c *Calculator) Calculate(queryTime float32) {

}
