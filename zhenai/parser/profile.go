package parser

import (
	"regexp"
	"strconv"

	"github.com/researchlab/crawler/engine"
	"github.com/researchlab/crawler/model"
)

var (
	ageRe    = regexp.MustCompile(`<div class="m-btn purple" data-v-ff544c08>([\d]+)岁</div>`)
	heightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-ff544c08>([\d]+)cm</div>`)
	weightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-ff544c08>([\d]+)kg</div>`)

	marriageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-ff544c08>(离异|未婚)</div>`)
	//	nameRe      = regexp.MustCompile(`<h1 class="nickName" data-v-35c72236>([^<]+)</h1>`)
	genderRe    = regexp.MustCompile(`genderString\":\"([^\"]+)\"`)
	incomeRe    = regexp.MustCompile(`<div class="m-btn purple" data-v-ff544c08>月收入:([^<]+)</div>`)
	educationRe = regexp.MustCompile(`educationString\":\"([^\"]+)\"`)
	hokouRe     = regexp.MustCompile(`<div class="m-btn pink" data-v-ff544c08>籍贯:([^<]+)</div>`)
	xinzuoRe    = regexp.MustCompile(`<div class="m-btn purple" data-v-ff544c08>((魔羯座|水瓶座|双鱼座|牡羊座|金牛座|双子座|巨蟹座|狮子座|天蝎座|射手座)\([\d]+.[\d]+-[\d]+.[\d]+\))</div>`)

	idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)
)

func ParseProfile(contents []byte, url, name string) engine.ParseResult {
	profile := model.Profile{}

	//	profile.Name = extractString(contents, nameRe)
	profile.Name = name
	profile.Gender = extractString(contents, genderRe)
	age, err := strconv.Atoi(
		extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}
	height, err := strconv.Atoi(
		extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}
	weight, err := strconv.Atoi(
		extractString(contents, weightRe))
	if err == nil {
		profile.Weight = weight
	}
	profile.Income = extractString(contents, incomeRe)
	profile.Marriage = extractString(contents, marriageRe)
	profile.Education = extractString(contents, educationRe)
	profile.Hokou = extractString(contents, hokouRe)
	profile.Xinzuo = extractString(contents, xinzuoRe)

	result := engine.ParseResult{
		Items: []engine.Item{
			engine.Item{
				Url:     url,
				Type:    "zhenai",
				Id:      extractString([]byte(url), idUrlRe),
				Payload: profile,
			},
		},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}
