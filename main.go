package main

import (
	ksuv "github.com/kooksee/ksuv/app"
	"flag"
)

var (
	cfg_path = flag.String("f", "config.ini", "配置文件的路径")
)

func main() {
	flag.Parse()

	//runtime.GOMAXPROCS(runtime.NumCPU())
	//defer log.Uninit(log.InitFile("./log/app.log"))
	//log.SetLevel(log.LvDEBUG)

	app := ksuv.GetApp()
	app.InitMiddleware()
	app.InitConfig(*cfg_path)
	app.InitLog()
	app.InitUrls()
	app.Run()
}



