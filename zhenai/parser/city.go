package parser

import (
	"crawler/engine"
	"regexp"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParseCity(contents []byte) engine.ParseResult {

	matches := profileRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			Parser:  NewProfileParser(name),
		})
	}

	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			Parser:engine.NewFuncParser(ParseCity, "ParseCity"),
		})
	}

	return result
}

type ProfileParser struct {
	userName string
}

func (p *ProfileParser) Parse(contents []byte) engine.ParseResult {
	return parseProfile(contents, p.userName)
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return "ProfileParser", p.userName
}

func NewProfileParser(name string) *ProfileParser {
	return &ProfileParser{userName:name}
}

