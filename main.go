package main

import (
	"github.com/miaoerduo/AwesomeReader/app"
	"github.com/miaoerduo/AwesomeReader/app/middleware"
)

func main() {
	app := app.EpubParser{
		EpubPath: "/Users/zhaoyu/Downloads/a.zip",
		MiddleWareList: []middleware.Middleware{
			&middleware.Dict{},
			&middleware.Menu{},
		}}
	app.Init()
	app.Dump("./out/")
}
