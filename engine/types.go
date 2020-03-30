package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url string
	Id string
	Type string
	Payload interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
