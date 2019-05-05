package parser

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("user_test_data.html")
	if err != nil {
		t.Error(err)
		t.Fail()
	} //
	result := ParseProfile(contents, "http://album.zhenai.com/u/108448601", "佐伊")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element; but was %v", result.Items)
	}
	//{"Name":"佐伊","Marriage":"未婚","Age":"27岁","Gender":"女士","Height":"172cm","Weight":"57kg","Income":"月收入:5-8千"
	// ,"Education":"大学本科","Occupation":"人事/行政","Hukou":"汉族","Xingzuo":"天秤座(09.23-10.22)","House":"租房","Car":"未买车","Workplace":"工作地:武汉蔡甸区","UserId":"108448601"}
	c, _ := json.Marshal(result.Items[0])
	t.Logf("%s", c)
}
