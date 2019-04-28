package parser

import (
	"crawler/engine"
	"crawler/model"
	"github.com/bitly/go-simplejson"
	"log"
	"regexp"
)

var profileRe = regexp.MustCompile(`<script>window.__INITIAL_STATE__=(.+);\(function`)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	match := profileRe.FindSubmatch(contents)
	result := engine.ParseResult{}
	if len(match) >= 2 {
		json := match[1]
		//fmt.Printf("%s\n", json)
		profile := parseJson(json)
		if profile == nil {
			return engine.ParseResult{}
		}
		profile.Name = name
		//bytes, _ := json2.Marshal(profile)
		//fmt.Println(string(bytes))
		//fmt.Println(profile)
		result.Items = append(result.Items, profile)
	}
	return result
}

func parseJson(json []byte) *model.Profile {
	res, err := simplejson.NewJson(json)
	if err != nil {
		log.Println("解析失败")
		return nil
	}
	objInfo := res.Get("objectInfo")

	infos, err := objInfo.Get("basicInfo").Array()

	//fmt.Printf("infos:%v,  %T\n", infos, infos)

	var profile model.Profile
	// 名字
	//nick, err := objInfo.Get("nickname").String()
	//if err == nil {
	//	profile.Name = nick
	//}
	gender, err := objInfo.Get("genderString").String()
	if err == nil {
		profile.Gender = gender
	}
	for k, v := range infos {
		if e, ok := v.(string); ok {
			switch k {
			case 0:
				profile.Marriage = e
			case 1:
				//年龄:47岁，我们可以设置int类型，所以可以通过另一个json字段来获取
				profile.Age = e
			case 2:
				profile.Xingzuo = e
			case 3:
				profile.Height = e
			case 4:
				profile.Weight = e
			case 5:
				profile.Workplace = e
			case 6:
				profile.Income = e
			case 7:
				profile.Occupation = e
			case 8:
				profile.Education = e
			}
		}
	}
	return &profile
}
