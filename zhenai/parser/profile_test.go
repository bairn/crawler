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

	result := ParseProfile(contents, "123")
	t.Logf("%#v\n", result)
}
