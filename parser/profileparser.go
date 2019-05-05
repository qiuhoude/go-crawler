package parser

import (
	"crawler/engine"
	"crawler/model"
	"github.com/bitly/go-simplejson"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var profileRe = regexp.MustCompile(`<script>window.__INITIAL_STATE__=(.+);\(function`)

func ParseProfile(contents []byte, url string, name string) engine.ParseResult {
	match := profileRe.FindSubmatch(contents)
	result := engine.ParseResult{}
	if len(match) >= 2 {
		jsonStr := match[1]
		//fmt.Printf("%s\n", jsonStr)
		profile := parseJson(jsonStr)
		if profile == nil {
			return result
		}
		//profile.Name = name
		//bytes, _ := json.Marshal(profile)
		//fmt.Println(string(bytes))
		//fmt.Println(profile)
		result.Items = append(result.Items, engine.Item{
			Url:     url,
			Id:      profile.UserId,
			Index:   "datint_profile",
			PayLoad: *profile,
		})
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
	nick, err := objInfo.Get("nickname").String()
	if err == nil {
		profile.Name = nick
	}
	// 性别
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

	infos2, err := res.Get("objectInfo").Get("detailInfo").Array()

	for _, v := range infos2 {
		/*
				"detailInfo": ["汉族", "籍贯:江苏宿迁", "体型:富线条美", "不吸烟", "不喝酒", "租房", "未买车", "没有小孩", "是否想要孩子:想要孩子", "何时结婚:认同闪婚"],
			   因为每个 每个用户的detailInfo数据不同，我们可以通过提取关键字来判断
		*/
		if e, ok := v.(string); ok {
			//fmt.Println(k, "--->", e)
			if strings.Contains(e, "族") {
				profile.Hukou = e
			} else if strings.Contains(e, "房") {
				profile.House = e
			} else if strings.Contains(e, "车") {
				profile.Car = e
			}
		}
	}

	//id
	if id, err := res.Get("objectInfo").Get("memberID").Int(); err == nil {
		profile.UserId = strconv.Itoa(id)
	}
	return &profile
}
