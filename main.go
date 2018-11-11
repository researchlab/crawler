package main

import (
	"github.com/researchlab/crawler/engine"
	"github.com/researchlab/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
