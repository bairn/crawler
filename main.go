package main

import (
	"github.com/bairn/crawler/engine"
	"github.com/bairn/crawler/persist"
	"github.com/bairn/crawler/scheduler"
	"github.com/bairn/crawler/zhenai/parser"
)

func main() {
	itemChan , err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:    itemChan,
		RequestProcessor:engine.Worker,
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/beijing",
		Parser: engine.NewFuncParser(parser.ParseCity, "ParseCity"),
	})
}
