package main

import (
	"crawler/engine"
	"crawler/persist"
	"crawler/scheduler"
	"crawler/zhenai/parser"
	"fmt"
)

func test(reader engine.ReadyNotifier) {
	fmt.Println(reader)
}

func main() {

	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})


	//e := engine.ConcurrentEngine{
	//	Scheduler:&scheduler.SimpleScheduler{},
	//	WorkerCount:10,
	//}
	//
	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})


	e := engine.ConcurrentEngine{
		Scheduler:&scheduler.QueuedScheduler{},
		WorkerCount:10,
		ItemChan:persist.ItemSaver(),
	}

	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/beijing",
		ParserFunc: parser.ParseCity,
	})

}