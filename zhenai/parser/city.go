package parser

import (
	"regexp"

	"github.com/researchlab/crawler/engine"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParseCity(
	contents []byte, url string) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}

	for _, m := range matches {
		//		result.Items = append(result.Items, "User "+name)
		url := string(m[1])
		result.Requests = append(
			result.Requests, engine.Request{
				Url:    url,
				Parser: NewProfileParser(string(m[2])),
			})
	}

	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests,
			engine.Request{
				Url:    string(m[1]),
				Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
			})
	}
	return result
}

type ProfileParser struct {
	userName string
}

func NewProfileParser(name string) *ProfileParser {
	return &ProfileParser{
		userName: name,
	}
}

func (p *ProfileParser) Parse(contents []byte, url string) engine.ParseResult {
	return parseProfile(contents, url, p.userName)
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return "ProfileParser", p.userName
}
