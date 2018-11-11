package main

import (
	"github.com/researchlab/crawler/engine"
	"github.com/researchlab/crawler/scheduler"
	"github.com/researchlab/crawler/zhenai/parser"
)

func main() {

	e := (&engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
	})
	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/shanghai",
		ParserFunc: parser.ParseCity,
	})
}
