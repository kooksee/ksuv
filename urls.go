package main

import (
	"github.com/gin-gonic/gin"
	"github.com/googollee/go-socket.io"
	"log"
)

func init_urls(r *gin.Engine) {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.On("connection", socketio_conn)
	server.On("error", socketio_error)

	r.GET("/", ping)
	r.GET("/ping", ping)
	r.POST("/command", ping)
	r.POST("/api/programs", programs_post)
	r.GET("/api/programs/:name", programs_get)
	r.GET("/api/programs/:name/ping", programs_status)
	r.GET("/api/programs/:name/status", programs_status)
	r.POST("/api/programs/:name/start", programs_start)
	r.POST("/api/programs/:name/stop", programs_stop)
	r.PUT("/api/programs/:name", programs_put)
	r.DELETE("/api/programs/:name", programs_delete)
	r.GET("/socket.io/", func(c *gin.Context) {
		server.ServeHTTP(c.Writer, c.Request)
	})
}
