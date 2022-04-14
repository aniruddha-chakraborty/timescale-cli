package services

import (
	"github.com/aniruddha-chakraborty/timescale-cli/interfaces"
	fifo "github.com/foize/go.fifo"
	"sync"
)

type Queue struct {
	queueStorage *fifo.Queue
}

func (q *Queue) Init() {
	q.queueStorage = fifo.NewQueue()
}

func ( q *Queue ) Start(input string) {
	var wg sync.WaitGroup
	for q.queueStorage.Len() > 0 {
		wg.Add(1)
		data := q.queueStorage.Next().(*interfaces.CsvStructure)
		go q.exec(data,wg)
	}
	wg.Wait()
}

func (q *Queue) exec(data *interfaces.CsvStructure, wg sync.WaitGroup) {

}

func (q *Queue) Insert(data *interfaces.CsvStructure) {
	q.queueStorage.Add(data)
}