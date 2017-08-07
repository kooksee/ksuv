package kapp

import (
	"sync"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"github.com/sirupsen/logrus"
	"os"
	"github.com/asdine/storm"
	"github.com/boltdb/bolt"
	"time"
)

var (
	once sync.Once
	instance *application
)

type application struct {
	services map[string]interface{}
	Cfg      *Configuration
	DB       *DB
	Log      *logrus.Logger
}

func (this *application) SetService(name string, service interface{}) {
	this.services[name] = service
}

func (this *application) GetService(name string) interface{} {
	return this.services[name]
}

func GetApp() *application {
	once.Do(func() {
		instance = &application{
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
	this.Cfg = cfg
}

// 初始化数据库
func (this *application)InitDB() {
	if !IsDir(this.Cfg.DbPath) {
		os.MkdirAll(this.Cfg.DbPath, 0600)
	}
	db, err := storm.Open(this.Cfg.DbPath+"/ksuv.db", storm.BoltOptions(0600, &bolt.Options{Timeout: 1 * time.Second}))
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
	if this.Cfg.Debug != "true" {
		logrus.SetFormatter(&logrus.JSONFormatter{})
		logrus.SetLevel(logrus.ErrorLevel)
		if file, err := os.OpenFile(this.Cfg.Log.Filepath, os.O_CREATE | os.O_WRONLY, 0666); err == nil {
			logrus.SetOutput(file)
		} else {
			panic("Failed to log to file, using default stderr")
		}
	} else {
		//log.SetFormatter(form)
		logrus.SetFormatter(&logrus.TextFormatter{})
		logrus.SetOutput(os.Stdout)
		logrus.SetLevel(logrus.DebugLevel)
	}
	this.Log = logrus.StandardLogger()



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
