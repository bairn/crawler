package main

import (
	"crawler/engine"
	"crawler/persist"
	"crawler/scheduler"
	"crawler/zhenai/parser"
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
