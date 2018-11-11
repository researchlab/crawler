package engine

import (
	"log"

	"github.com/researchlab/crawler/fetcher"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (p *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	p.Scheduler.Run()

	for i := 0; i < p.WorkerCount; i++ {
		createWorker(p.Scheduler.WorkerChan(), out, p.Scheduler)
	}

	for _, r := range seeds {
		p.Scheduler.Submit(r)
	}
	count := 0
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item#%d: %v\n", count, item)
			count++
		}
		for _, request := range result.Requests {
			p.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body), nil
}
