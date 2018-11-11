package scheduler

import "github.com/researchlab/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (p *SimpleScheduler) Submit(r engine.Request) {
	go func() { p.workerChan <- r }()
}

func (p *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	p.workerChan = c
}
