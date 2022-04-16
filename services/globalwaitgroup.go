package services

import "sync"

type GlobalWaitGroup struct {
	wg *sync.WaitGroup
}

func (g *GlobalWaitGroup)  Init() {
	g.wg = &sync.WaitGroup{}
}

func (g *GlobalWaitGroup)  Instance() *sync.WaitGroup {
	return g.wg
}

