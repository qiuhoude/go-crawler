package parser

import (
	"crawler/fetcher"
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		t.Fail()
	}
	parseResult := ParseCityList(contents)
	const resultSize = 470
	if len(parseResult.Requests) != 470 {
		t.Errorf("result should have %d requests, but had %d", resultSize, len(parseResult.Requests))
	}
	if len(parseResult.Items) != 470 {
		t.Errorf("result should have %d Items, but had %d", resultSize, len(parseResult.Items))
	}
}

func TestParseCityList4Json(t *testing.T) {
	contents, _ := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	parseResult := ParseCityList4Json(contents)
	const resultSize = 470
	if len(parseResult.Requests) != 470 {
		t.Errorf("result should have %d requests, but had %d", resultSize, len(parseResult.Requests))
	}
}
