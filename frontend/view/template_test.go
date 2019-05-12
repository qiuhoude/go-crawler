package view

import (
	"crawler/engine"
	"crawler/frontend/model"
	model2 "crawler/model"
	"os"
	"testing"
)

func TestTemplate(t *testing.T) {

	view := CreateSearchResultView("template.html")

	file, err := os.Create("template_test.html")

	page := model.SearchResult{
		Hits:  123,
		Start: 0,
	}

	profile := model2.Profile{
		Name:       "林YY",
		Marriage:   "未婚",
		Age:        "26岁",
		Xingzuo:    "魔羯座(12.22-01.19)",
		Height:     "165cm",
		Weight:     "50kg",
		Income:     "月收入:5-8千",
		Occupation: "职业技术教师",
		Education:  "高中及以下",
		Gender:     "女",
		Hukou:      "湖北",
	}
	item := engine.Item{
		Url:     "http://album.zhenai.com/u/1214814888",
		Index:   "datint_profile",
		Id:      "1214814888",
		PayLoad: profile,
	}

	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}

	err = view.Render(file, page)
	if err != nil {
		t.Log(err)
	}

}
