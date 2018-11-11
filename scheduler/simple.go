package scheduler

import "github.com/researchlab/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (p *SimpleScheduler) Submit(r engine.Request) {
	go func() { p.workerChan <- r }()
}

func (p *SimpleScheduler) Run() {
	p.workerChan = make(chan engine.Request)
}

func (p *SimpleScheduler) WorkerChan() chan engine.Request {
	return p.workerChan
}

func (p *SimpleScheduler) WorkerReady(chan engine.Request) {}
