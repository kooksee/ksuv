package app

import (
	"sync"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"fmt"
	"github.com/manucorporat/stats"
	log "github.com/sirupsen/logrus"
	"os"
	"github.com/asdine/storm"
	"github.com/boltdb/bolt"
	"time"
)

var (
	ips = stats.New()
)

var (
	once sync.Once
	instance *application
)

type application struct {
	services map[string]interface{}
	gin      *gin.Engine
	cfg      *Configuration
	DB       *DB
	Log      *log.Logger
}

func (this *application) SetService(name string, service interface{}) {
	this.services[name] = service
}

func (this *application) GetService(name string) interface{} {
	return this.services[name]
}

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

func (this *application)InitMiddleware() {
	if this.cfg.Debug == "true" {
		this.gin.Use(gin.Logger())
	} else {
	}

	this.gin.Use(rateLimit, gin.Recovery())

}

func (this *application)Run() {
	this.gin.Run()
}

func GetApp() *application {
	once.Do(func() {
		instance = &application{
			gin:gin.New(),
			services: make(map[string]interface{}),
		}
	})

	return instance
}

func (this *application)InitConfig(cfg_path string) {
	data, err := ioutil.ReadFile(cfg_path)
	if err != nil {
		panic(err)

	}

	cfg := &Configuration{}
	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		panic(err)
	}
	this.cfg = cfg
}

// 初始化数据库
func (this *application)InitDB() {

	db, err := storm.Open(this.cfg.DbPath, storm.BoltOptions(0600, &bolt.Options{Timeout: 1 * time.Second}))
	if err != nil {
		panic(err)
	}
	app_db := &DB{
		Scripts:db.From("scripts"),
		Programs : db.From("programs"),
		Logs :db.From("logs"),
		Sessions : db.From("sessions"),
		Status : db.From("status"),
		DB:db,
	}
	this.DB = app_db
}

func (this *application)InitLog() {
	if this.cfg.Debug != "true" {
		log.SetFormatter(&log.JSONFormatter{})
		log.SetLevel(log.ErrorLevel)
		if file, err := os.OpenFile(this.cfg.Log.Filepath, os.O_CREATE | os.O_WRONLY, 0666); err == nil {
			log.SetOutput(file)
		} else {
			panic("Failed to log to file, using default stderr")
		}
	} else {
		//log.SetFormatter(form)
		log.SetFormatter(&log.TextFormatter{})
		log.SetOutput(os.Stdout)
		log.SetLevel(log.DebugLevel)
	}
	this.Log = log.StandardLogger()



	// You could set this to any `io.Writer` such as a file
	// file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY, 0666)
	// if err == nil {
	//  log.Out = file
	// } else {
	//  log.Info("Failed to log to file, using default stderr")
	// }

	// Log as JSON instead of the default ASCII formatter.
	//log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	//log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	//log.SetLevel(log.InfoLevel)
}
