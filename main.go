package main

import (
	"github.com/researchlab/crawler/engine"
	"github.com/researchlab/crawler/scheduler"
	"github.com/researchlab/crawler/zhenai/parser"
)

func main() {

	(&engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
	}).Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
