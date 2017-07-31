package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"log"
	"os/signal"
	"syscall"
	"fmt"
	"time"
	"github.com/googollee/go-socket.io"
)

func main() {

	r := gin.New()
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.On("connection", socketio_conn)
	server.On("error", socketio_error)



	// Global middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

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

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 100)


	// `signal.Notify` registers the given channel to
	// receive notifications of the specified signals.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	// This goroutine executes a blocking receive for
	// signals. When it gets one it'll print it out
	// and then notify the program that it can finish.
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		//done <- true
		os.Exit(0)
	}()

	go func() {
		ticker := time.NewTicker(time.Second * 2)
		defer ticker.Stop()
		done <- true

		for {
			select {
			case v1 := <-done:
				fmt.Println("job.....", v1)
			case <-ticker.C:
				done <- false
				fmt.Println("job.....")
			}
		}
	}()

	r.Run() // listen and serve on 0.0.0.0:8080


}



