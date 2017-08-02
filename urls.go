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

	// 服务操作
	// 添加服务
	r.POST("/api/programs", programs_post)

	// 修改服务信息
	r.PUT("/api/programs/:name", programs_put)

	// 根据服务名称获取服务信息
	r.GET("/api/programs/:name", programs_get)

	// 删除服务
	r.DELETE("/api/programs/:name", programs_delete)

	// 获取服务的存活状态
	r.GET("/api/programs/:name/ping", programs_status)

	// 获取服务的进程,IO,CPU使用等信息
	r.GET("/api/programs/:name/status", programs_status)

	// 创建启动服务
	r.POST("/api/programs/:name", programs_start)

	// 获得服务的运行结果
	r.GET("/api/programs/:name/:id", programs_start)

	// 查看当前服务的运行状态
	r.GET("/api/programs/:name/:id/ping", programs_start)

	// 查看当前服务的运行监控信息
	r.GET("/api/programs/:name/:id/status", programs_start)

	// 暂停服务
	r.DELETE("/api/programs/:name/:id", programs_stop)


	// 服务日志操作
	// 根据时间戳获取日志信息
	r.GET("/api/programs/:name/logs", programs_stop)

	// 获取该服务的日志存储状况
	r.GET("/api/programs/:name/logs/status", programs_stop)

	// 根据天删除之前的日志信息
	r.DELETE("/api/programs/:name/logs", programs_stop)


	// 脚本管理管理
	// 添加shell信息
	r.POST("/api/scripts", programs_stop)

	// 修改脚本信息
	r.PUT("/api/scripts/:name", programs_stop)

	// 获取脚本信息
	r.GET("/api/scripts/:name", programs_stop)

	// 删除脚本
	r.DELETE("/api/scripts/:name", programs_stop)

	// 创建执行脚本session
	r.POST("/api/scripts/:name", programs_stop)

	// 启动创建的脚本session
	r.POST("/api/scripts/:name/:id", programs_stop)

	// 查看脚本时候还在运行
	r.GET("/api/scripts/:name/:id/ping", programs_stop)

	// 获得脚本执行状态
	r.GET("/api/scripts/:name/:id/status", programs_stop)

	// 获得脚本执行的结果
	r.POST("/api/scripts/:name/:id", programs_stop)

	// 删除该session
	r.DELETE("/api/scripts/:name/:id", programs_stop)

	r.GET("/socket.io/", func(c *gin.Context) {
		server.ServeHTTP(c.Writer, c.Request)
	})
}
