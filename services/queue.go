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

func ( q *Queue ) Start(input string) {
	var wg sync.WaitGroup
	for q.queueStorage.Len() > 0 {
		wg.Add(1)
		data := q.queueStorage.Next().(*interfaces.CsvStructure)
		go q.exec(data,wg,q.Calculator)
	}
	wg.Wait()
}

func (q *Queue) exec(data *interfaces.CsvStructure, wg sync.WaitGroup, calculator *Calculator) {
	start := time.Now()
	time.Sleep(1 * time.Second)
	end   := time.Now()
	d := end.UnixNano() - start.UnixNano()
	calculator.Calculate(d)
	defer wg.Done()
}

func (q *Queue) Insert(data *interfaces.CsvStructure) {
	q.queueStorage.Add(data)
}