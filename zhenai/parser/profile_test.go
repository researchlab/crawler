package parser

import (
	"io/ioutil"
	"testing"

	"github.com/researchlab/crawler/model"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile(
		"profile_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "蜗牛漫步")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 "+"element; but was %v", result.Items)
	}

	profile := result.Items[0].(model.Profile)

	expected := model.Profile{
		Name:      "蜗牛漫步",
		Gender:    "男士",
		Age:       29,
		Height:    175,
		Weight:    59,
		Income:    "5-8千",
		Marriage:  "未婚",
		Education: "中专",
		Hokou:     "四川绵阳",
		Xinzuo:    "魔羯座(12.22-01.19)",
	}

	if profile.Name != expected.Name ||
		profile.Gender != expected.Gender ||
		profile.Age != expected.Age ||
		profile.Height != expected.Height ||
		profile.Weight != expected.Weight ||
		profile.Income != expected.Income ||
		profile.Marriage != expected.Marriage ||
		profile.Education != expected.Education ||
		profile.Hokou != expected.Hokou ||
		profile.Xinzuo != expected.Xinzuo {
		t.Errorf("profile is not expected %v", profile)
	}
}
