package persist

import (
	"context"
	"crawler/engine"
	"crawler/model"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"testing"
)

// 测试es获取
func TestGetItem(t *testing.T) {
	//从ElasticSearch中获取，根据id
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(EsUrl))
	resp, err := client.Get().
		Index("datint_profile").
		Id("1409115162").
		Do(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s", resp.Source) //打印

}
func TestItemSaver(t *testing.T) {
	// "basicInfo": ["未婚", "26岁", "魔羯座(12.22-01.19)", "165cm", "50kg", "工作地:苏州相城区", "月收入:5-8千", "职业技术教师", "高中及以下"],
	// &{Index:datint_profile Type:_doc Id:XOA7gWoBl8ghTlcLsNx0 Version:1 Result:created Shards:0xc0000bd6e0 SeqNo:2 PrimaryTerm:1 Status:0 ForcedRefresh:false}--- PASS: TestItemSaver (0.16s)
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
	item := engine.Item{
		Url:     "http://album.zhenai.com/u/1214814888",
		Index:   "datint_profile",
		Id:      "1214814888",
		PayLoad: profile,
	}

	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(EsUrl))
	id, err := Save(client, item)

	t.Logf("id = %s", id)
	if err != nil {
		t.Fatal(err)
	}
	//从ElasticSearch中获取，根据id
	resp, err := client.Get().
		Index(item.Index).
		Id(item.Id).
		Do(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s", resp.Source) //打印

	//反序列化
	var actual engine.Item
	err = json.Unmarshal([]byte(resp.Source), &actual)

	if err != nil {
		t.Fatal(err)
	}

	obj, _ := model.FromJsonObj(actual.PayLoad)
	actual.PayLoad = obj

	//断言
	if actual != item {
		t.Errorf("got %v; expected %v", actual, profile)
	}

}
