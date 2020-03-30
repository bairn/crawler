package parser

import (
	"crawler/engine"
	"crawler/model"
	"github.com/bitly/go-simplejson"
	"log"
	"regexp"
	"strconv"
)

var re = regexp.MustCompile(`<script>window.__INITIAL_STATE__=(.+);\(function`)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	match := re.FindSubmatch(contents)
	result := engine.ParseResult{}
	if len(match) >= 2 {
		json := match[1]

		item := parseJson(json, name)

		result.Items = append(result.Items, item)
	}

	return result
}



func parseJson(json []byte, name string) engine.Item {
	res, err := simplejson.NewJson(json)
	if err != nil {
		log.Fatalln("解析json失败")
	}

	infos, err := res.Get("objectInfo").Get("basicInfo").Array()
	if err != nil {
		log.Fatalln(err)
	}

	var item engine.Item
	var profile model.Profile

	if len(infos) != 9 {
		return engine.Item{}
	}

	profile.Name = name
	for k, v := range infos {
		if e, ok := v.(string); ok {
			switch k {
			case 0:
				profile.Marriage = e
			case 1:
				profile.Age = e
			case 2:
				profile.Xingzuo = e
			case 3:
				profile.Height = e
			case 4:
				profile.Weight = e
			case 6:
				profile.Income = e
			case 7:
				profile.Occupation = e
			case 8:
				profile.Education = e
			}
		}
	}


	memberId, err := res.Get("objectInfo").Get("memberID").Int()
	if err != nil {
		panic(err)
	}

	memberIdString := strconv.Itoa(memberId)
	item = engine.Item{
		Url:     "https://album.zhenai.com/u/" + memberIdString,
		Id:      memberIdString,
		Type:    "zhenai",
		Payload: profile,
	}

	return item
}
