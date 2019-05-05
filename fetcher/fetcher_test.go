package fetcher

import (
	"testing"
)

func TestFetcher(t *testing.T) {

	contents, err := Fetch("http://album.zhenai.com/u/108448601")
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	t.Logf("%s", string(contents))
}
