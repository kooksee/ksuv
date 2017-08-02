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

	// 服务资源操作
	// 添加多个服务资源
	r.POST("/api/programs", programs_post)

	// 修改多个服务的信息
	r.PUT("/api/programs", programs_put)

	// 修改单个服务的信息
	r.PUT("/api/programs/:name", programs_put)

	// 获取多个服务信息
	r.GET("/api/programs", programs_get)

	// 根据服务名称获取单个服务信息
	r.GET("/api/programs/:name", programs_get)

	// 删除多个服务资源
	r.DELETE("/api/programs", programs_delete)

	// 删除单个服务资源
	r.DELETE("/api/programs/:name", programs_delete)

	// 服务操作
	// 创建启动服务的session
	r.POST("/api/programs/:name/sessions", programs_service_post)

	// 获得服务的运行结果
	r.GET("/api/programs/sessions", programs_service_get)

	// 获得服务的运行结果
	r.GET("/api/programs/sessions/:id", programs_service_get)

	// 查看所有服务的运行状态
	r.GET("/api/programs/sessions/ping", programs_service_ping)

	// 查看当前服务的运行状态
	r.GET("/api/programs/sessions/:id/ping", programs_service_ping)

	// 查看所有服务的运行监控信息
	r.GET("/api/programs/sessions/status", programs_service_status)

	// 暂停服务
	r.DELETE("/api/programs/sessions", programs_service_delete)

	// 暂停服务
	r.DELETE("/api/programs/sessions/:id", programs_service_delete)

	// 服务日志操作
	// 获取日志信息
	r.GET("/api/logs", programs_log_get)

	// 根据服务名称获取日志信息
	r.GET("/api/logs/:name", programs_log_get)

	// 获得具体信息
	r.GET("/api/logs/:name/:id", programs_log_get)

	// 获取该服务的日志存储状况
	r.GET("/api/logs/status", programs_log_status)

	// 根据天删除之前的日志信息
	r.DELETE("/api/logs", programs_log_delete)

	// 根据天删除之前的日志信息
	r.DELETE("/api/logs/:date", programs_log_delete)


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
