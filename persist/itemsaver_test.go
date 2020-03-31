package persist

import (
	"context"
	"crawler/engine"
	"crawler/model"
	"encoding/json"
	"gopkg.in/olivere/elastic.v5"
	"testing"
)

func TestItemSaver(t *testing.T) {

	expected := engine.Item{
		Url:     "https://album.zhenai.com/u/1682657696",
		Id:      "1682657696",
		Type:    "zhenai",
		Payload: model.Profile{
			Name:       "紫色玫瑰",
			Marriage:   "未婚",
			Age:        "28岁",
			Gender:     "",
			Height:     "161cm",
			Weight:     "49kg",
			Income:     "月收入:2-5万",
			Education:  "大学本科",
			Occupation: "计算机/互联网",
			Hukou:      "",
			Xingzuo:    "魔羯座(12.22-01.19)",
			House:      "",
			Car:        "",
		},
	}

	client, err := elastic.NewClient(
		elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	const index = "dating_test"
	// Fetch saved item
	err = save(index, client, expected)
	if err != nil {
		panic(err)
	}

	//TODO:Try to start up elastic search

	resp, err := client.Get().
		Index("dating_profile").
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	//t.Logf("%s", *resp.Source)

	var actual engine.Item
	//json.Unmarshal([]byte(*resp.Source), &actual)
	json.Unmarshal(*resp.Source, &actual)

	actualProfile , _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	if actual != expected {
		t.Errorf("got %v;expected %v", actual, expected)
	}
}