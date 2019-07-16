package parser

import (
	"io/ioutil"
	"testing"
)

func TestCityParser(t *testing.T) {
	contents, err := ioutil.ReadFile("city_test_data.html")
	if err != nil {
		t.Fail()
	}
	parseResult := ParseCity(contents, "")
	for _, c := range parseResult.Items {
		t.Logf("Get user nick is %q", c)
	}
}
