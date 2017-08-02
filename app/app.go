package app

import (
	"sync"
	"github.com/gin-gonic/gin"
	"github.com/mkideal/log"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

var (
	once sync.Once
	instance *application
)

type config struct {

}

type application struct {
	services map[string]interface{}
	gin      *gin.Engine
	cfg      Configuration
}

func NewApp() *application {
	return &application{
		gin:gin.New(),
	}

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
		instance = &application{services: make(map[string]interface{})}
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

func (this *application)InitDb(map[string]string) {
}

func (this *application)InitLog() {
	defer log.Uninit(log.InitFileAndConsole(this.cfg.Log.Filepath, log.LvERROR))
}
