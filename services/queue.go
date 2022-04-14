package services

import (
	"github.com/aniruddha-chakraborty/timescale-cli/interfaces"
	fifo "github.com/foize/go.fifo"
	"sync"
	"time"
)

type Queue struct {
	queueStorage *fifo.Queue
	Calculator *Calculator
}

func (q *Queue) Init() {
	q.queueStorage = fifo.NewQueue()
}

func ( q *Queue ) Start() {
	var wg sync.WaitGroup
	startTime := time.Now()
	totalInQueue := int64(q.queueStorage.Len())
	for q.queueStorage.Len() > 0 {
		wg.Add(1)
		data := q.queueStorage.Next().(*interfaces.CsvStructure)
		go q.exec(data,wg,q.Calculator,startTime,totalInQueue)
	}
	wg.Wait()
}

func (q *Queue) exec(data *interfaces.CsvStructure, wg sync.WaitGroup, calculator *Calculator,startTime time.Time,totalInQueue int64) {
	start := time.Now()
	time.Sleep(1 * time.Second)
	end   := time.Now()
	d := end.UnixNano() - start.UnixNano()
	calculator.Calculate(d,startTime,totalInQueue)
	defer wg.Done()
}

func (q *Queue) Insert(data *interfaces.CsvStructure) {
	q.queueStorage.Add(data)
}