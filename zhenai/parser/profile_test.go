package parser

import (
	"crawler/fetcher"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := fetcher.Fetch("https://album.zhenai.com/u/1674105583")
	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "安静的雪")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element;but was %v", result.Items)
	}

	t.Logf("%#v\n", result)
}
