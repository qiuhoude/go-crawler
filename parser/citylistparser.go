package parser

import (
	"crawler/engine"
	"github.com/bitly/go-simplejson"
	"log"
	"regexp"
)

var cityListRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^<]+)</a>`)
var cityList4JsonRe = regexp.MustCompile(`<script>window.__INITIAL_STATE__=(.+);\(function`)

/*
解析城市列表
*/
func ParseCityList(contents []byte) engine.ParseResult {
	all := cityListRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, c := range all {
		//result.Items = append(result.Items, string(c[2])) //城市名称
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(c[1]),
			ParserFunc: ParseCity,
		})
	}
	return result
}

func ParseCityList4Json(contents []byte) engine.ParseResult {
	jsonb := cityList4JsonRe.FindSubmatch(contents)
	all := parseJsonCityList(jsonb[1])
	result := engine.ParseResult{}
	for _, c := range all {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(c[1]),
			ParserFunc: ParseCity,
		})
	}
	return result
}

func parseJsonCityList(jsonb []byte) [][]string {
	res, err := simplejson.NewJson(jsonb)
	if err != nil {
		log.Println("解析json失败。。")
	}
	infos, _ := res.Get("cityListData").Get("cityData").Array()
	//infos是一个切片，里面的类型是interface{}

	var dataList [][]string
	//所以我们遍历这个切片，里面使用断言来判断类型
	for _, v := range infos {
		if each_map, ok := v.(map[string]interface{}); ok {
			//fmt.Println(each_map)
			map2 := each_map["cityList"]
			for _, v2 := range map2.([]interface{}) {
				if data, ok := v2.(map[string]interface{}); ok {
					var datas []string
					cityName := data["linkContent"].(string)
					cityUrl := data["linkURL"].(string)
					datas = append(datas, cityName, cityUrl)
					dataList = append(dataList, datas)
				}
			}
		}
	}
	return dataList
}
