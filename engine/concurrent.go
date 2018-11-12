package engine

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan Item
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
		if isDuplicate(r.Url) {
			continue
		}
		p.Scheduler.Submit(r)
	}
	for {
		result := <-out
		for _, item := range result.Items {
			//log.Printf("Got item: %v\n", item)
			go func() { p.ItemChan <- item }()
		}
		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				continue
			}
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

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}
