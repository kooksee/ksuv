package main

import (
	ksuv "github.com/kooksee/ksuv/app"
	"flag"
	//"github.com/mkideal/log"
	//"github.com/gin-gonic/gin"
)

var (
	cfg_path = flag.String("f", "config.yml", "配置文件的路径")
)

func main() {
	//defer log.Uninit(log.InitWithLogger(gin.Logger()))
	//defer log.Uninit(log.InitConsole(log.LvINFO))

	flag.Parse()

	//runtime.GOMAXPROCS(runtime.NumCPU())
	//defer log.Uninit(log.InitFile("./log/app.log"))
	//log.SetLevel(log.LvDEBUG)
	//gin.SetMode(gin.ReleaseMode)


	app := ksuv.GetApp()
	app.InitConfig(*cfg_path)

	app.InitMiddleware()
	app.InitLog()
	app.InitUrls()
	app.Run()
}



