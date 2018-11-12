package persist

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/olivere/elastic"
	"github.com/researchlab/crawler/engine"
	"github.com/researchlab/crawler/model"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/91162834",
		Type: "zhenai",
		Id:   "91162834",
		Payload: model.Profile{
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
		},
	}

	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		t.Fatal(err)
	}

	const index = "dating_test"
	// Save expected item
	err = Save(client, index, expected)
	if err != nil {
		t.Error(err)
	}

	// Fetch saved item
	resp, err := client.Get().Index(index).
		Type(expected.Type).Id(expected.Id).Do(context.Background())

	if err != nil {
		t.Error(err)
	}
	t.Logf("%s", *resp.Source)

	var actual engine.Item
	err = json.Unmarshal([]byte(*resp.Source), &actual)
	if err != nil {
		t.Error(err)
	}

	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile
	//verfiy result
	if actual != expected {
		t.Errorf("got %v; expected got %v", actual, expected)
	}
}
