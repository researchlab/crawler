package parser

import (
	"regexp"

	"github.com/researchlab/crawler/engine"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte, url string) engine.ParseResult {
	re, _ := regexp.Compile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		//		result.Items = append(
		//			result.Items, string(m[2]))
		result.Requests = append(
			result.Requests, engine.Request{
				Url:    string(m[1]),
				Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
			})
	}
	return result
}
