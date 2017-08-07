package main

import (
	"github.com/kooksee/ksuv/kapp"
	"flag"
	//"github.com/mkideal/log"
	//"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/manucorporat/stats"
)

var (
	cfg_path = flag.String("f", "config.yml", "配置文件的路径")
	ips = stats.New()
)

func rateLimit(c *gin.Context) {
	ip := c.ClientIP()
	value := int(ips.Add(ip, 1))
	if value % 50 == 0 {
		fmt.Printf("ip: %s, count: %d\n", ip, value)
	}
	if value >= 200 {
		if value % 200 == 0 {
			fmt.Println("ip blocked")
		}
		c.Abort()
		c.String(503, "you were automatically banned :)")
	}
}
func main() {
	//defer log.Uninit(log.InitFile("./log/app.log"))
	//runtime.GOMAXPROCS(runtime.NumCPU())
	//log.SetLevel(log.LvDEBUG)
	//gin.SetMode(gin.ReleaseMode)

	flag.Parse()

	k := kapp.GetApp()
	k.InitConfig(*cfg_path)
	k.InitLog()
	k.InitDB()

	app := gin.New()
	if k.Cfg.Debug == "true" {
		app.Use(gin.Logger())
	} else {
		app.Use(rateLimit, gin.Recovery())
	}

	InitUrls(app)
	app.Run()
}



