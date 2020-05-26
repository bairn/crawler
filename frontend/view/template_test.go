package view

import (
	"github.com/bairn/crawler/engine"
	"github.com/bairn/crawler/frontend/model"
	common "github.com/bairn/crawler/model"
	"os"
	"testing"
)

func TestSearchResultView_Render(t *testing.T)  {
	veiw := CreateSearchResultView("template.html")

	out, err := os.Create("template.test.html")
	if err != nil {
		panic(err)
	}

	page := model.SearchResult{}
	page.Hits = 124
	item := engine.Item{
		Url:     "https://album.zhenai.com/u/1682657696",
		Id:      "1682657696",
		Type:    "zhenai",
		Payload: common.Profile{
			Name:       "心瘾",
			Marriage:   "离异",
			Age:        "34岁",
			Gender:     "女",
			Height:     "160",
			Weight:     "50kg",
			Income:     "12000",
			Education:  "本科",
			Occupation: "",
			Hukou:      "",
			Xingzuo:    "",
			House:      "",
			Car:        "",
		},
	}
	for i:=0;i<10;i++ {
		page.Items = append(page.Items, item)
	}

	err = veiw.Render(out, page)
	if err != nil {
		panic(err)
	}
}
