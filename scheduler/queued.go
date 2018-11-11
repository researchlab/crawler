package scheduler

import "github.com/researchlab/crawler/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (p *QueuedScheduler) Submit(r engine.Request) {
	p.requestChan <- r
}

func (p *QueuedScheduler) WorkerReady(w chan engine.Request) {
	p.workerChan <- w
}

func (p *QueuedScheduler) ConfigureMasterWorkerChan(chan engine.Request) {}

func (p *QueuedScheduler) Run() {
	p.workerChan = make(chan chan engine.Request)
	p.requestChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request

		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}

			select {
			case r := <-p.requestChan:
				requestQ = append(requestQ, r)
			case w := <-p.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}
