package persist

import (
	"crawler/model"
	"testing"
)

func TestItemSaver(t *testing.T) {
	// "basicInfo": ["未婚", "26岁", "魔羯座(12.22-01.19)", "165cm", "50kg", "工作地:苏州相城区", "月收入:5-8千", "职业技术教师", "高中及以下"],

	profile := model.Profile{
		Name:       "林YY",
		Marriage:   "未婚",
		Age:        "26岁",
		Xingzuo:    "魔羯座(12.22-01.19)",
		Height:     "165cm",
		Weight:     "50kg",
		Income:     "月收入:5-8千",
		Occupation: "职业技术教师",
		Education:  "高中及以下",
	}
	save(profile)
}
