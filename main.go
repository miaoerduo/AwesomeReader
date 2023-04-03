package main

import "github.com/miaoerduo/AwesomeReader/app"

func main() {
	app := app.EpubParser{EpubPath: "/Users/zhaoyu/Desktop/GGS.epub"}
	app.Init()
	app.Dump("./out/")
}
