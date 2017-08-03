package app

import (
	"sync"
	"github.com/gin-gonic/gin"
	"github.com/mkideal/log"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"github.com/kooksee/ksuv/db"
)

var (
	once sync.Once
	instance *application
)

type application struct {
	services map[string]interface{}
	gin      *gin.Engine
	cfg      Configuration
	DB       *db.DB
}

func (this *application) SetService(name string, service interface{}) {
	this.services[name] = service
}

func (this *application) GetService(name string) interface{} {
	return this.services[name]
}

func (this *application)InitMiddleware() {
	this.gin.Use(gin.Logger())
	this.gin.Use(gin.Recovery())
}

func (this *application)Run() {
	this.gin.Run()
}

func GetApp() *application {
	once.Do(func() {
		instance = &application{
			gin:gin.New(),
			services: make(map[string]interface{}),
			cfg: &Configuration{},
		}
	})

	return instance
}

func (this *application)InitConfig(cfg_path string) {
	data, err := ioutil.ReadFile(cfg_path)
	if err != nil {
		panic(err)

	}

	err = yaml.Unmarshal(data, &this.cfg)
	if err != nil {
		panic(err)
	}
}

func (this *application)InitDB() {
	this.DB.InitDB(this.cfg.DbPath)
}

func (this *application)InitLog() {
	defer log.Uninit(log.InitFileAndConsole(this.cfg.Log.Filepath, log.LvERROR))
}
