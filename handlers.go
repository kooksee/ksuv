package main

import (
	"github.com/gin-gonic/gin"
	"github.com/googollee/go-socket.io"
	"fmt"
	"log"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func programs_post(c *gin.Context) {
}

func programs_delete(c *gin.Context) {
	program_name := c.Params.ByName("name")
	println(program_name)
	c.JSON(200, gin.H{
		"message": program_name,
	})
}

func programs_put(c *gin.Context) {
	program_name := c.Params.ByName("name")
	println(program_name)
	c.JSON(200, gin.H{
		"message": program_name,
	})
}

func programs_get(c *gin.Context) {
	program_name := c.Params.ByName("name")
	println(program_name)
	c.JSON(200, gin.H{
		"message": program_name,
	})
}

func programs_status(c *gin.Context) {
	program_name := c.Params.ByName("name")
	println(program_name)
	c.JSON(200, gin.H{
		"message": program_name,
	})
}

func programs_start(c *gin.Context) {
	program_name := c.Params.ByName("name")
	println(program_name)
	c.JSON(200, gin.H{
		"message": program_name,
	})
}

func programs_stop(c *gin.Context) {
	program_name := c.Params.ByName("name")
	println(program_name)
	c.JSON(200, gin.H{
		"message": program_name,
	})
}

func socketio_conn(so socketio.Socket) {
	log.Println("on connection")
	so.Join("chat")
	so.On("chat message", func(msg string) {
		m := make(map[string]interface{})
		m["a"] = "你好"
		e := so.Emit("cn1111", m)
		//这个没有问题
		fmt.Println("\n\n")

		b := make(map[string]string)
		b["u-a"] = "中文内容" //这个不能是中文
		m["b-c"] = b
		e = so.Emit("cn2222", m)
		log.Println(e)

		log.Println("emit:", so.Emit("chat message", msg))
		so.BroadcastTo("chat", "chat message", msg)
	})
	// Socket.io acknowledgement example
	// The return type may vary depending on whether you will return
	// For this example it is "string" type
	so.On("chat message with ack", func(msg string) string {
		return msg
	})
	so.On("disconnection", func() {
		log.Println("on disconnect")
	})
}

func socketio_error(so socketio.Socket, err error) {
	log.Println("error:", err)
}